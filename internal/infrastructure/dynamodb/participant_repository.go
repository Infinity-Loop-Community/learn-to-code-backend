package dynamodb

import (
	"context"
	"encoding/json"
	"fmt"
	"learn-to-code/internal/domain/eventsource"
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/domain/quiz/participant/event"
	"learn-to-code/internal/infrastructure/config"
	"reflect"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

type MarshalFunc func(v interface{}) ([]byte, error)
type UnmarshalFunc func(data []byte, v interface{}) error

type EventPo struct {
	AggregateID string    `dynamodbav:"aggregate_id"`
	Type        string    `dynamodbav:"type"`
	Version     uint      `dynamodbav:"version"`
	Payload     string    `dynamodbav:"payload"`
	CreatedAt   time.Time `dynamodbav:"created_at"`
}

type ParticipantRepository struct {
	dbClient     *dynamodb.Client
	ctx          context.Context
	serializer   MarshalFunc
	deserializer UnmarshalFunc
	tableName    string
}

func NewDynamoDbParticipantRepository(ctx context.Context, environment config.Environment, client *dynamodb.Client) *ParticipantRepository {

	tableName := fmt.Sprintf("%s_events", environment)

	return &ParticipantRepository{
		dbClient:     client,
		ctx:          ctx,
		serializer:   json.Marshal,
		deserializer: json.Unmarshal,
		tableName:    tableName,
	}
}

func (r ParticipantRepository) StoreEvents(participantID string, events []eventsource.Event) error {
	for _, e := range events {
		err := r.appendEvent(participantID, e)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r ParticipantRepository) appendEvent(participantID string, e eventsource.Event) error {
	serializedEvent, err := r.serializer(e)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item: map[string]types.AttributeValue{
			"aggregate_id": &types.AttributeValueMemberS{Value: participantID},
			"version":      &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", e.GetVersion())},
			"type":         &types.AttributeValueMemberS{Value: reflect.TypeOf(e).Name()},
			"payload":      &types.AttributeValueMemberS{Value: string(serializedEvent)},
			"created_at":   &types.AttributeValueMemberS{Value: e.GetCreatedAt().Format(time.RFC3339)},
		},
	}
	_, err = r.dbClient.PutItem(r.ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (r ParticipantRepository) FindEventsByID(pariticipantID string) ([]eventsource.Event, error) {
	input := &dynamodb.QueryInput{
		TableName: &r.tableName,
		KeyConditions: map[string]types.Condition{
			"aggregate_id": {
				ComparisonOperator: types.ComparisonOperatorEq,
				AttributeValueList: []types.AttributeValue{
					&types.AttributeValueMemberS{Value: pariticipantID},
				},
			},
		},
	}

	output, err := r.dbClient.Query(r.ctx, input)
	if err != nil {
		return []eventsource.Event{}, err
	}

	var events []eventsource.Event

	for _, outputItem := range output.Items {
		deserializedEvent, deserializeError := r.outputItemToEvent(outputItem)

		if deserializeError != nil {
			return nil, deserializeError
		}

		events = append(events, deserializedEvent)
	}

	return events, nil

}

func (r ParticipantRepository) FindOrCreateByID(id string) (participant.Participant, error) {
	input := &dynamodb.QueryInput{
		TableName: &r.tableName,
		KeyConditions: map[string]types.Condition{
			"aggregate_id": {
				ComparisonOperator: types.ComparisonOperatorEq,
				AttributeValueList: []types.AttributeValue{
					&types.AttributeValueMemberS{Value: id},
				},
			},
		},
	}

	output, err := r.dbClient.Query(r.ctx, input)
	if err != nil {
		return participant.Participant{}, err
	}

	if len(output.Items) == 0 {
		return participant.NewParticipant(id)
	}

	var events []eventsource.Event

	for _, outputItem := range output.Items {

		deserializedEvent, deserializeError := r.outputItemToEvent(outputItem)

		if deserializeError != nil {
			return participant.Participant{}, deserializeError
		}
		events = append(events, deserializedEvent)

	}

	p, newFromEventsErr := participant.NewFromEvents(events, true)

	return p, newFromEventsErr
}

func (r ParticipantRepository) outputItemToEvent(outputItem map[string]types.AttributeValue) (eventsource.Event, error) {
	eventPo := EventPo{}
	err := attributevalue.UnmarshalMap(outputItem, &eventPo)
	if err != nil {
		return nil, err
	}

	var deserializeError error
	var deserializedEvent eventsource.Event

	switch eventPo.Type {

	case event.ParticipantCreatedTypeName:
		joinedQuizEvent := &event.ParticipantCreated{}

		deserializeError = r.deserializer([]byte(eventPo.Payload), joinedQuizEvent)
		deserializedEvent = *joinedQuizEvent

	case event.StartedQuizTypeName:
		startedQuiz := &event.StartedQuiz{}

		deserializeError = r.deserializer([]byte(eventPo.Payload), startedQuiz)

		deserializedEvent = *startedQuiz

	case event.SelectedAnswerTypeName:
		e := &event.SelectedAnswer{}

		deserializeError = r.deserializer([]byte(eventPo.Payload), e)

		deserializedEvent = *e

	case event.FinishedQuizTypeName:
		finishedQuiz := &event.FinishedQuiz{}

		deserializeError = r.deserializer([]byte(eventPo.Payload), finishedQuiz)

		deserializedEvent = *finishedQuiz

	default:
		panic(fmt.Errorf("unknown type '%s' while reading persisted events", eventPo.Type))
	}
	return deserializedEvent, deserializeError
}

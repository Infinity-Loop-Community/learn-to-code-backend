package dynamodb

import (
	"context"
	"encoding/json"
	"fmt"
	"learn-to-code/internal/domain/eventsource"
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/infrastructure/config"
	"reflect"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type MarshalFunc func(v interface{}) ([]byte, error)

type ParticipantRepository struct {
	dbClient            *dynamodb.Client
	eventPODeserializer *EventPODeserializer
	ctx                 context.Context
	serializer          MarshalFunc
	tableName           string
}

func NewDynamoDbParticipantRepository(ctx context.Context, environment config.Environment, client *dynamodb.Client, eventPODeserializer *EventPODeserializer) *ParticipantRepository {

	tableName := fmt.Sprintf("%s_events", environment)

	return &ParticipantRepository{
		dbClient:            client,
		ctx:                 ctx,
		serializer:          json.Marshal,
		tableName:           tableName,
		eventPODeserializer: eventPODeserializer,
	}
}

func (r *ParticipantRepository) FindEventsByParticipantID(participantID string) ([]eventsource.Event, error) {
	output, err := r.findEventsByParticipantID(participantID)
	if err != nil {
		return []eventsource.Event{}, err
	}

	events, err := r.queryOutputToEvents(output)
	if err != nil {
		return []eventsource.Event{}, err
	}

	return events, nil
}

func (r *ParticipantRepository) FindOrCreateByID(id string) (participant.Participant, error) {
	output, err := r.findEventsByParticipantID(id)
	if err != nil {
		return participant.Participant{}, err
	}

	if len(output.Items) == 0 {
		return participant.NewParticipant(id)
	}

	events, err := r.queryOutputToEvents(output)
	if err != nil {
		return participant.Participant{}, err
	}

	p, newFromEventsErr := participant.NewFromEvents(events, true)

	return p, newFromEventsErr
}

func (r *ParticipantRepository) StoreEvents(participantID string, events []eventsource.Event) error {
	for _, e := range events {
		err := r.appendEvent(participantID, e)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ParticipantRepository) appendEvent(participantID string, e eventsource.Event) error {
	serializedEvent, err := r.serializer(e)
	if err != nil {
		return err
	}

	err = r.putEventForParticipantID(participantID, e, serializedEvent, err)
	if err != nil {
		return err
	}

	return nil
}

func (r *ParticipantRepository) putEventForParticipantID(participantID string, e eventsource.Event, serializedEvent []byte, err error) error {
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
	return err
}

func (r *ParticipantRepository) queryOutputToEvents(output *dynamodb.QueryOutput) ([]eventsource.Event, error) {
	var events []eventsource.Event

	for _, outputItem := range output.Items {

		deserializedEvent, deserializeError := r.eventPODeserializer.outputItemToEvent(outputItem)

		if deserializeError != nil {
			return nil, deserializeError
		}
		events = append(events, deserializedEvent)

	}
	return events, nil
}

func (r *ParticipantRepository) findEventsByParticipantID(id string) (*dynamodb.QueryOutput, error) {
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

	return r.dbClient.Query(r.ctx, input)
}

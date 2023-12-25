package dynamodb

import (
	"encoding/json"
	"fmt"
	"learn-to-code/internal/domain/eventsource"
	"learn-to-code/internal/domain/quiz/participant/event"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UnmarshalFunc func(data []byte, v interface{}) error

type EventPODeserializer struct {
	deserializer UnmarshalFunc
}

func NewEventPODeserializer() *EventPODeserializer {
	return &EventPODeserializer{
		deserializer: json.Unmarshal,
	}
}

func (r EventPODeserializer) outputItemToEvent(outputItem map[string]types.AttributeValue) (eventsource.Event, error) {
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

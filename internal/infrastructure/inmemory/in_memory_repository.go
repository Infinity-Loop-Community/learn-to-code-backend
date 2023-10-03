package inmemory

import (
	"encoding/json"
	"fmt"
	"hello-world/internal/domain/eventsource"
	. "hello-world/internal/domain/quiz/participant"
	"hello-world/internal/domain/quiz/participant/event"
	"reflect"
	"time"
)

type MarshalFunc func(v interface{}) ([]byte, error)
type UnmarshalFunc func(data []byte, v interface{}) error

type ParticipantRepository struct {
	data         map[string][]EventPo
	serializer   MarshalFunc
	deserializer UnmarshalFunc
}

type EventPo struct {
	AggregateID string
	Type        string
	Version     uint
	Payload     []byte
	CreatedAt   time.Time
}

func NewParticipantRepository() *ParticipantRepository {
	return &ParticipantRepository{
		data: make(map[string][]EventPo),

		serializer:   json.Marshal,
		deserializer: json.Unmarshal,
	}
}

func (r *ParticipantRepository) AppendEvent(id string, e eventsource.Event) error {
	serializedEvent, err := r.serializer(e)
	if err != nil {
		return err
	}

	po := EventPo{
		AggregateID: e.GetId(),
		Version:     e.GetVersion(),
		Type:        reflect.TypeOf(e).Name(),
		Payload:     serializedEvent,
		CreatedAt:   e.GetCreatedAt(),
	}
	r.data[id] = append(r.data[id], po)
	return nil
}

func (r *ParticipantRepository) FindById(id string) (Participant, error) {
	eventPos, ok := r.data[id]

	if ok != true {
		return Participant{}, ErrNotFound
	}

	var events []eventsource.Event

	for _, po := range eventPos {

		switch po.Type {

		case event.JoinedQuizTypeName:
			joinedQuizEvent := &event.JoinedQuiz{}

			err := r.deserializer(po.Payload, joinedQuizEvent)
			if err != nil {
				return Participant{}, err
			}

			events = append(events, *joinedQuizEvent)
		case event.FinishedQuizTypeName:
			finishedQuiz := &event.FinishedQuiz{}

			err := r.deserializer(po.Payload, finishedQuiz)
			if err != nil {
				return Participant{}, err
			}

			events = append(events, *finishedQuiz)

		case event.StartedQuizTypeName:
			startedQuiz := &event.StartedQuiz{}

			err := r.deserializer(po.Payload, startedQuiz)
			if err != nil {
				return Participant{}, err
			}

			events = append(events, *startedQuiz)

		default:
			panic(fmt.Errorf("unknown type '%s' while reading persisted events", po.Type))
		}

	}

	p, err := NewFromEvents(events)

	return p, err
}

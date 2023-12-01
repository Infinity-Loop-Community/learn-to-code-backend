package inmemory

import (
	"encoding/json"
	"fmt"
	"learn-to-code/internal/domain/eventsource"
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/domain/quiz/participant/event"
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
	ID          string
	AggregateID string
	Type        string
	Version     uint
	Payload     []byte
	CreatedAt   time.Time
}

var GlobalParticipantData = make(map[string][]EventPo)

func NewParticipantRepository() *ParticipantRepository {
	return &ParticipantRepository{
		data: GlobalParticipantData,

		serializer:   json.Marshal,
		deserializer: json.Unmarshal,
	}
}

func (r *ParticipantRepository) StoreEvents(id string, events []eventsource.Event) error {
	for _, e := range events {
		err := r.storeEvent(id, e)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ParticipantRepository) storeEvent(id string, e eventsource.Event) error {
	serializedEvent, err := r.serializer(e)
	if err != nil {
		return err
	}

	po := EventPo{
		AggregateID: e.GetAggregateID(),
		Version:     e.GetVersion(),
		Type:        reflect.TypeOf(e).Name(),
		Payload:     serializedEvent,
		CreatedAt:   e.GetCreatedAt(),
	}
	r.data[id] = append(r.data[id], po)
	return nil
}

func (r *ParticipantRepository) FindOrCreateByID(id string) (participant.Participant, error) {
	eventPos, ok := r.data[id]

	if !ok {
		newParticipant, err := participant.NewParticipant(id)
		if err != nil {
			return participant.Participant{}, err
		}

		events := newParticipant.GetNewEventsAndUpdatePersistedVersion()
		r.StoreEvents(id, events)

		eventPos, ok = r.data[id]
		if !ok {
			return participant.Participant{}, fmt.Errorf("could not fetch data from before created list of events")
		}
	}

	var events []eventsource.Event

	for _, po := range eventPos {

		switch po.Type {

		case event.ParticipantCreatedTypeName:
			participantCreatedEvent := &event.ParticipantCreated{}

			err := r.deserializer(po.Payload, participantCreatedEvent)
			if err != nil {
				return participant.Participant{}, err
			}

			events = append(events, *participantCreatedEvent)

		case event.StartedQuizTypeName:
			startedQuiz := &event.StartedQuiz{}

			err := r.deserializer(po.Payload, startedQuiz)
			if err != nil {
				return participant.Participant{}, err
			}

			events = append(events, *startedQuiz)

		case event.SelectedAnswerTypeName:
			event := &event.SelectedAnswer{}

			err := r.deserializer(po.Payload, event)
			if err != nil {
				return participant.Participant{}, err
			}

			events = append(events, *event)

		case event.FinishedQuizTypeName:
			finishedQuiz := &event.FinishedQuiz{}

			err := r.deserializer(po.Payload, finishedQuiz)
			if err != nil {
				return participant.Participant{}, err
			}

			events = append(events, *finishedQuiz)
		default:
			panic(fmt.Errorf("unknown type '%s' while reading persisted events", po.Type))
		}

	}

	p, err := participant.NewFromEvents(events, true)

	return p, err
}

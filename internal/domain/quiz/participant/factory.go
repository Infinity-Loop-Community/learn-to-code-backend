package participant

import (
	"learn-to-code/internal/domain/eventsource"
	"learn-to-code/internal/domain/quiz/participant/event"
	"learn-to-code/internal/infrastructure/go/util/uuid"
	"time"
)

func New() (Participant, error) {
	return NewWithId(uuid.MustNewRandomAsString())
}

func NewWithId(id string) (Participant, error) {
	joinedQuizEvent := event.ParticipantCreated{
		EventBase: eventsource.EventBase{
			Id:        id,
			Version:   0,
			CreatedAt: time.Now(),
		},
	}
	return NewFromEvents([]eventsource.Event{joinedQuizEvent})
}

func NewFromEvents(events []eventsource.Event) (Participant, error) {

	p := Participant{}

	for _, e := range events {

		err := p.apply(e)

		if err != nil {
			return Participant{}, err
		}
	}

	return p, nil
}

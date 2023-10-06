package participant

import (
	"hello-world/internal/domain/eventsource"
	"hello-world/internal/domain/quiz/participant/event"
	"hello-world/internal/infrastructure/go/util/uuid"
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

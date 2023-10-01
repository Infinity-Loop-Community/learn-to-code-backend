package participant

import (
	"hello-world/internal/domain/quiz/participant/event"
	"hello-world/internal/infrastructure/go/util/uuid"
	"time"
)

func New() (Participant, error) {
	joinedQuizEvent := event.JoinedQuiz{
		Id:   uuid.MustNewRandomAsString(),
		Time: time.Now(),
	}
	return NewFromEvents([]event.Event{joinedQuizEvent})
}

func NewFromEvents(events []event.Event) (Participant, error) {
	p := Participant{}

	for _, e := range events {
		err := p.apply(e)

		if err != nil {
			return Participant{}, err
		}
	}

	return p, nil
}

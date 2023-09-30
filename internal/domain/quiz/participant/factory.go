package participant

import (
	"hello-world/internal/domain/quiz/participant/event"
	"hello-world/internal/infrastructure/go/util/uuid"
)

func New() Participant {
	return Participant{
		id:      uuid.MustNewRandomAsString(),
		Quizzes: nil,
		events:  nil,
	}
}

func NewFromEvents(id string, events []event.Event) (Participant, error) {
	p := Participant{
		id:      id,
		Quizzes: nil,
		events:  nil,
	}

	for _, e := range events {
		err := p.apply(e)

		if err != nil {
			return Participant{}, err
		}
	}

	return p, nil
}

package participant

import (
	"learn-to-code/internal/domain/eventsource"
	"learn-to-code/internal/domain/quiz/participant/event"
	"learn-to-code/internal/infrastructure/go/util/uuid"
	"time"
)

func New() (Participant, error) {
	return NewParticipant(uuid.MustNewRandomAsString())
}

func NewParticipant(id string) (Participant, error) {
	participantCreated := event.ParticipantCreated{
		EventBase: eventsource.EventBase{
			AggregateID: id,
			Version:     0,
			CreatedAt:   time.Now(),
		},
	}
	return NewFromEvents([]eventsource.Event{participantCreated}, false)
}

func NewFromEvents(events []eventsource.Event, isPersisted bool) (Participant, error) {

	p := Participant{
		quizzes: map[string][]*activeQuiz{},
	}

	for _, e := range events {

		err := p.apply(e, isPersisted)

		if err != nil {
			return Participant{}, err
		}
	}

	return p, nil
}

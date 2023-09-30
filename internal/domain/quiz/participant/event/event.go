package event

import "hello-world/internal/domain/quiz/participant"

type Event interface {
	CheckIfApplicable(p *participant.Participant) error
}

package event

import (
	"fmt"
	"hello-world/internal/domain/quiz/participant"
)

type StartedQuiz struct {
	Id string
}

func (e StartedQuiz) CheckIfApplicable(p *participant.Participant) error {
	for _, quiz := range p.Quizzes {
		if quiz.Id == e.Id && quiz.IsOngoing() {
			return fmt.Errorf("quiz '%s' already started and not finished", quiz.Id)
		}
	}
	return nil
}

package event

import (
	"fmt"
	"hello-world/internal/domain/quiz/participant"
)

type FinishedQuiz struct {
	Id string
}

func (e FinishedQuiz) CheckIfApplicable(p *participant.Participant) error {
	quizFound := false
	for _, quiz := range p.Quizzes {
		if quiz.Id == e.Id {
			quizFound = true
			break
		}
	}

	if quizFound == false {
		return fmt.Errorf("no started quiz found with Id '%s'", e.Id)
	}
	return nil
}

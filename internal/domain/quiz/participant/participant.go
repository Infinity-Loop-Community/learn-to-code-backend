package participant

import (
	"fmt"
	"hello-world/internal/domain/quiz/participant/event"
	"reflect"
	"time"
)

type Participant struct {
	id      string
	joined  time.Time
	Quizzes []activeQuiz

	events []event.Event
}

func (p *Participant) apply(e event.Event) error {

	switch e.(type) {

	case event.JoinedQuiz:
		p.id, p.joined = e.(event.JoinedQuiz).Id, e.(event.JoinedQuiz).Time

	case event.StartedQuiz:
		p.appendToQuizList(e.(event.StartedQuiz).Id)

	case event.FinishedQuiz:
		p.setQuizCompleted(e.(event.FinishedQuiz).Id)

	default:
		panic(fmt.Sprintf("unknown event type %s", reflect.TypeOf(e)))
	}

	return nil
}

func (p *Participant) setQuizCompleted(finishedQuizId string) {
	for i, quiz := range p.Quizzes {
		if quiz.Id == finishedQuizId {
			p.Quizzes[i].completed = true
			break
		}
	}
}

func (p *Participant) appendToQuizList(eventQuizId string) {
	p.Quizzes = append(p.Quizzes, activeQuiz{
		Id:              eventQuizId,
		providedAnswers: nil,
		completed:       false,
	})
}

func (p *Participant) GetId() string {
	return p.id
}

func (p *Participant) GetStartedQuizCount() int {
	return len(p.Quizzes)
}

func (p *Participant) GetFinishedQuizCount() int {
	finishedQuizzes := 0

	for _, quiz := range p.Quizzes {
		if quiz.completed {
			finishedQuizzes++
		}
	}

	return finishedQuizzes
}

func (p *Participant) StartQuiz(id string) (event.StartedQuiz, error) {

	quiz, err := p.ensureQuizNotStarted(id)
	if err != nil {
		return quiz, err
	}

	var startedQuizEvent = event.StartedQuiz{
		Id: id,
	}

	err = p.apply(startedQuizEvent)

	return startedQuizEvent, err
}

func (p *Participant) ensureQuizNotStarted(id string) (event.StartedQuiz, error) {
	for _, quiz := range p.Quizzes {
		if quiz.Id == id && quiz.IsOngoing() {
			return event.StartedQuiz{}, fmt.Errorf("quiz '%s' already started and not finished", quiz.Id)
		}
	}
	return event.StartedQuiz{}, nil
}

func (p *Participant) FinishQuiz(id string) (event.FinishedQuiz, error) {
	var foundQuiz *activeQuiz

	for _, quiz := range p.Quizzes {

		if quiz.Id == id && quiz.completed != true {
			foundQuiz = &quiz
			break
		}
	}

	if foundQuiz == nil {
		return event.FinishedQuiz{}, fmt.Errorf("quiz not found")
	}

	finishedQuizEvent := event.FinishedQuiz{
		Id: id,
	}

	err := p.apply(finishedQuizEvent)

	return finishedQuizEvent, err
}

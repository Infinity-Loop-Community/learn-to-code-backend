package participant

import (
	"fmt"
	"hello-world/internal/domain/quiz/participant/event"
	"reflect"
)

type Participant struct {
	id      string
	quizzes []activeQuiz

	events []event.Event
}

func (p *Participant) apply(e event.Event) error {
	switch e.(type) {

	case event.StartedQuiz:
		err2 := p.handleStartedQuizEvent(e)
		if err2 != nil {
			return err2
		}

	case event.FinishedQuiz:
		err := p.handleFinishedQuizEvent(e)
		if err != nil {
			return err
		}

	default:
		panic(fmt.Sprintf("unknown event type %s", reflect.TypeOf(e)))
	}

	return nil
}

func (p *Participant) handleFinishedQuizEvent(e event.Event) error {
	finishedQuizEvent := e.(event.FinishedQuiz)

	err := p.allowsApplyFinishedQuizEvent(finishedQuizEvent.Id)
	if err != nil {
		return err
	}

	p.setQuizCompleted(finishedQuizEvent.Id)

	return nil
}

func (p *Participant) setQuizCompleted(finishedQuizId string) {
	for i, quiz := range p.quizzes {
		if quiz.id == finishedQuizId {
			p.quizzes[i].completed = true
			break
		}
	}
}

func (p *Participant) allowsApplyFinishedQuizEvent(finishedQuizEventId string) error {
	quizFound := false
	for _, quiz := range p.quizzes {
		if quiz.id == finishedQuizEventId {
			quizFound = true
			break
		}
	}

	if quizFound == false {
		return fmt.Errorf("no started quiz found with id '%s'", finishedQuizEventId)
	}
	return nil
}

func (p *Participant) handleStartedQuizEvent(e event.Event) error {
	eventQuizId := e.(event.StartedQuiz).Id

	err := p.allowsApplyStartedQuizEvent(eventQuizId)
	if err != nil {
		return err
	}

	p.appendToQuizList(eventQuizId)
	return nil
}

func (p *Participant) allowsApplyStartedQuizEvent(eventQuizId string) error {
	for _, quiz := range p.quizzes {
		if quiz.id == eventQuizId && quiz.isOngoing() {
			return fmt.Errorf("quiz '%s' already started and not finished", quiz.id)
		}
	}
	return nil
}

func (p *Participant) appendToQuizList(eventQuizId string) {
	p.quizzes = append(p.quizzes, activeQuiz{
		id:              eventQuizId,
		providedAnswers: nil,
		completed:       false,
	})
}

func (p *Participant) GetId() string {
	return p.id
}

func (p *Participant) GetStartedQuizCount() int {
	return len(p.quizzes)
}

func (p *Participant) GetFinishedQuizCount() int {
	finishedQuizzes := 0

	for _, quiz := range p.quizzes {
		if quiz.completed {
			finishedQuizzes++
		}
	}

	return finishedQuizzes
}

func (p *Participant) StartQuiz(id string) (event.StartedQuiz, error) {

	var startedQuizEvent = event.StartedQuiz{
		Id: id,
	}

	err := p.apply(startedQuizEvent)

	return startedQuizEvent, err
}

func (p *Participant) FinishQuiz(id string) (event.FinishedQuiz, error) {
	var foundQuiz *activeQuiz

	for _, quiz := range p.quizzes {

		if quiz.id == id && quiz.completed != true {
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

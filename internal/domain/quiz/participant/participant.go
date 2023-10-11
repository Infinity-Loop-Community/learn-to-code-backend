package participant

import (
	"fmt"
	"learn-to-code/internal/domain/eventsource"
	"learn-to-code/internal/domain/quiz/participant/event"
	"reflect"
	"time"
)

type Participant struct {
	id      string
	Quizzes []activeQuiz

	eventsource.AggregateRoot
}

func (p *Participant) apply(e eventsource.Event) error {

	switch eventType := e.(type) {

	case event.ParticipantCreated:
		p.id = eventType.ID

	case event.StartedQuiz:
		p.appendToQuizList(eventType.ID)

	case event.FinishedQuiz:
		p.setQuizCompleted(eventType.ID)

	default:
		panic(fmt.Sprintf("unknown event type %s", reflect.TypeOf(e)))
	}

	p.CurrentVersion++
	p.Events = append(p.Events, e)

	return nil
}

func (p *Participant) setQuizCompleted(finishedQuizID string) {
	for i, quiz := range p.Quizzes {
		if quiz.ID == finishedQuizID {
			p.Quizzes[i].completed = true
			break
		}
	}
}

func (p *Participant) appendToQuizList(eventQuizID string) {
	p.Quizzes = append(p.Quizzes, activeQuiz{
		ID:              eventQuizID,
		providedAnswers: nil,
		completed:       false,
	})
}

func (p *Participant) GetID() string {
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

func (p *Participant) StartQuiz(id string) error {

	err := p.ensureQuizNotStarted(id)
	if err != nil {
		return err
	}

	var startedQuizEvent = event.StartedQuiz{
		EventBase: p.createEventBaseEvent(id),
	}

	err = p.apply(startedQuizEvent)

	return err
}

func (p *Participant) ensureQuizNotStarted(id string) error {
	for _, quiz := range p.Quizzes {
		if quiz.ID == id && quiz.IsOngoing() {
			return fmt.Errorf("quiz '%s' already started and not finished", quiz.ID)
		}
	}
	return nil
}

func (p *Participant) FinishQuiz(id string) error {
	var foundQuiz *activeQuiz

	for _, quiz := range p.Quizzes {

		if quiz.ID == id && !quiz.completed {
			foundQuiz = &quiz
			break
		}
	}

	if foundQuiz == nil {
		return fmt.Errorf("quiz not found")
	}

	finishedQuizEvent := event.FinishedQuiz{
		EventBase: p.createEventBaseEvent(id),
	}

	err := p.apply(finishedQuizEvent)

	return err
}

func (p *Participant) createEventBaseEvent(id string) eventsource.EventBase {
	return eventsource.EventBase{
		ID:        id,
		Version:   p.CurrentVersion,
		CreatedAt: time.Now(),
	}
}

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
	quizzes map[string]*activeQuiz

	eventsource.AggregateRoot
}

func (p *Participant) apply(eventToApply eventsource.Event) error {

	var err error

	switch e := eventToApply.(type) {

	case event.ParticipantCreated:
		p.id = e.GetAggregateID()

	case event.StartedQuiz:
		err := p.ensureQuizNotStarted(e.QuizID)
		if err != nil {
			return err
		}

		p.quizzes[e.QuizID] = &activeQuiz{
			ID:              e.QuizID,
			providedAnswers: nil,
			completed:       false,
		}

	case event.SelectedAnswer:
		quiz, ok := p.quizzes[e.QuizID]
		if !ok {
			return fmt.Errorf("quiz %v not found", e.QuizID)
		}

		if quiz.completed {
			return fmt.Errorf("can not selected an answer for quiz %v that is already completed", e.QuizID)
		}

		quiz.providedAnswers = append(quiz.providedAnswers, ProvidedAnswer{
			QuestionID: e.QuestionID,
			AnswerID:   e.AnswerID,
		})

	case event.FinishedQuiz:
		_, ok := p.quizzes[e.QuizID]

		if !ok {
			return fmt.Errorf("quiz %v not found", e.QuizID)
		}

		quiz, ok := p.quizzes[e.QuizID]
		if !ok {
			return fmt.Errorf("Quiz not started %v, hence it can not be completed", e.QuizID)
		}

		if quiz.completed {
			return fmt.Errorf("Quiz %v already finished", e.QuizID)
		}

		quiz.completed = true
		if err != nil {
			return err
		}

	default:
		panic(fmt.Sprintf("unknown event type %s", reflect.TypeOf(eventToApply)))
	}

	p.CurrentVersion++
	p.Events = append(p.Events, eventToApply)

	return nil
}

func (p *Participant) ensureQuizNotStarted(id string) error {
	for _, quiz := range p.quizzes {
		if quiz.ID == id && quiz.IsOngoing() {
			return fmt.Errorf("quiz '%s' already started and not finished", quiz.ID)
		}
	}
	return nil
}

func (p *Participant) StartQuiz(quizID string) error {

	var startedQuizEvent = event.StartedQuiz{
		EventBase: p.createEventBaseEvent(),
		QuizID:    quizID,
	}

	err := p.apply(startedQuizEvent)

	return err
}

func (p *Participant) FinishQuiz(quizID string) error {
	finishedQuizEvent := event.FinishedQuiz{
		EventBase: p.createEventBaseEvent(),
		QuizID:    quizID,
	}

	err := p.apply(finishedQuizEvent)

	return err
}

func (p *Participant) SelectQuizAnswer(quizID string, questionID string, answerID string) error {
	selectedAnswerEvent := event.SelectedAnswer{
		QuizID:     quizID,
		QuestionID: questionID,
		AnswerID:   answerID,
		EventBase:  p.createEventBaseEvent(),
	}

	err := p.apply(selectedAnswerEvent)

	return err
}

func (p *Participant) createEventBaseEvent() eventsource.EventBase {
	return eventsource.EventBase{
		AggregateID: p.id,
		Version:     p.CurrentVersion,
		CreatedAt:   time.Now(),
	}
}

func (p *Participant) GetID() string {
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

func (p *Participant) GetActiveQuizAnswers(quizID string) ([]ProvidedAnswer, error) {
	quiz, ok := p.quizzes[quizID]
	if !ok {
		return nil, fmt.Errorf("quiz %v not found", quizID)
	}

	return quiz.providedAnswers, nil
}

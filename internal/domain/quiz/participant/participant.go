package participant

import (
	"fmt"
	"learn-to-code/internal/domain/eventsource"
	"learn-to-code/internal/domain/quiz/participant/event"
	"reflect"
	"time"
)

type Participant struct {
	// id is the unique identifier for the Participant, distinguishing each user within
	// the system.
	id string

	// quizAttempts holds information about active quizAttempts associated with the participant. This
	// includes data about the quizAttempts the participant is currently engaged with and their progress.
	quizAttempts map[string][]*quizAttempt

	eventsource.AggregateRoot
}

func (p *Participant) apply(eventToApply eventsource.Event, isPersisted bool) error {

	switch ev := eventToApply.(type) {

	case event.ParticipantCreated:
		p.id = ev.GetAggregateID()

	case event.StartedQuiz:
		err := p.ensureQuizNotStarted(ev.QuizID)
		if err != nil {
			return err
		}

		p.quizAttempts[ev.QuizID] = append(p.quizAttempts[ev.QuizID], &quizAttempt{
			ID:                        ev.QuizID,
			providedAnswers:           nil,
			requiredQuestionsAnswered: ev.RequiredQuestionsAnswered,
			completed:                 false,
		})

	case event.SelectedAnswer:
		quizAttempts, ok := p.quizAttempts[ev.QuizID]
		if !ok {
			return fmt.Errorf("quizAttempt %v not found", ev.QuizID)
		}
		quizAttemptCount := len(quizAttempts)
		quiz := quizAttempts[quizAttemptCount-1]

		if quiz.completed {
			return fmt.Errorf("can not selected an answer for quizAttempt %v that is already completed", ev.QuizID)
		}

		quiz.providedAnswers = append(quiz.providedAnswers, ProvidedAnswer{
			QuestionID: ev.QuestionID,
			AnswerID:   ev.AnswerID,
		})

	case event.FinishedQuiz:
		quizAttempt, err := p.getCurrentQuizAttempt(ev)
		if err != nil {
			return err
		}

		err = quizAttempt.checkFinishAttemptValidity()
		if err != nil {
			return err
		}

		quizAttempt.completed = true

	default:
		panic(fmt.Sprintf("unknown event type %s", reflect.TypeOf(eventToApply)))
	}

	p.AppendEvent(eventToApply, isPersisted)

	return nil
}

func (p *Participant) getCurrentQuizAttempt(ev event.FinishedQuiz) (*quizAttempt, error) {
	quizAttempts, ok := p.quizAttempts[ev.QuizID]
	if !ok {
		return nil, fmt.Errorf("quizAttempt %v not found", ev.QuizID)
	}

	quizAttemptCount := len(quizAttempts)
	quizAttempt := quizAttempts[quizAttemptCount-1]
	return quizAttempt, nil
}

func (p *Participant) ensureQuizNotStarted(id string) error {
	for _, quizAttempts := range p.quizAttempts {

		quizAttemptCount := len(quizAttempts)
		quiz := quizAttempts[quizAttemptCount-1]

		if quiz.ID == id && quiz.IsOngoing() {
			return fmt.Errorf("quiz '%s' already started and not finished", quiz.ID)
		}
	}
	return nil
}

func (p *Participant) StartQuiz(quizID string, requiredQuestionsAnswered []string) error {

	var startedQuizEvent = event.StartedQuiz{
		EventBase:                 p.createEventBaseEvent(),
		QuizID:                    quizID,
		RequiredQuestionsAnswered: requiredQuestionsAnswered,
	}

	err := p.apply(startedQuizEvent, false)

	return err
}

func (p *Participant) FinishQuiz(quizID string) error {
	finishedQuizEvent := event.FinishedQuiz{
		EventBase: p.createEventBaseEvent(),
		QuizID:    quizID,
	}

	err := p.apply(finishedQuizEvent, false)

	return err
}

func (p *Participant) SelectQuizAnswer(quizID string, questionID string, answerID string) error {
	selectedAnswerEvent := event.SelectedAnswer{
		QuizID:     quizID,
		QuestionID: questionID,
		AnswerID:   answerID,
		EventBase:  p.createEventBaseEvent(),
	}

	err := p.apply(selectedAnswerEvent, false)

	return err
}

func (p *Participant) createEventBaseEvent() eventsource.EventBase {
	return eventsource.EventBase{
		AggregateID: p.id,
		Version:     p.GetCurrentVersion(),
		CreatedAt:   time.Now(),
	}
}

func (p *Participant) GetID() string {
	return p.id
}

func (p *Participant) GetStartedQuizCount() int {
	return len(p.quizAttempts)
}

func (p *Participant) GetFinishedQuizCount() int {
	finishedQuizzes := 0

	for _, quizAttempts := range p.quizAttempts {

		quizAttemptCount := len(quizAttempts)
		quiz := quizAttempts[quizAttemptCount-1]

		if quiz.completed {
			finishedQuizzes++
		}
	}

	return finishedQuizzes
}

func (p *Participant) GetActiveQuizAnswers(quizID string) ([]ProvidedAnswer, error) {
	quizAttempts, ok := p.quizAttempts[quizID]
	if !ok {
		return nil, fmt.Errorf("quiz %v not found", quizID)
	}
	quizAttemptCount := len(quizAttempts)
	quiz := quizAttempts[quizAttemptCount-1]

	return quiz.providedAnswers, nil
}

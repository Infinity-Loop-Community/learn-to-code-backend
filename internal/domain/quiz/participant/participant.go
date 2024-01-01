package participant

import (
	"fmt"
	"learn-to-code/internal/domain/eventsource"
	"learn-to-code/internal/domain/quiz/participant/event"
	"reflect"
	"strconv"
	"time"
)

type Participant struct {
	id           string
	quizAttempts map[string][]*quizAttempt

	eventsource.AggregateRoot
}

func (p *Participant) apply(eventToApply eventsource.Event, isPersisted bool) error {

	switch e := eventToApply.(type) {

	case event.ParticipantCreated:
		p.id = e.GetAggregateID()

	case event.StartedQuiz:
		err := p.ensureQuizNotStarted(e.QuizID)
		if err != nil {
			return err
		}

		p.quizAttempts[e.QuizID] = append(p.quizAttempts[e.QuizID], &quizAttempt{
			QuizID:                    e.QuizID,
			providedAnswers:           nil,
			requiredQuestionsAnswered: e.RequiredQuestionsAnswered,
			completed:                 false,
		})

	case event.SelectedAnswer:
		quizAttempts, ok := p.quizAttempts[e.QuizID]
		if !ok {
			return fmt.Errorf("lastQuizAttempt %v not found", e.QuizID)
		}
		quizAttemptCount := len(quizAttempts)
		quiz := quizAttempts[quizAttemptCount-1]

		if quiz.completed {
			return fmt.Errorf("can not selected an answer for lastQuizAttempt %v that is already completed", e.QuizID)
		}

		quiz.providedAnswers = append(quiz.providedAnswers, ProvidedAnswer{
			QuestionID: e.QuestionID,
			AnswerID:   e.AnswerID,
			IsCorrect:  e.IsCorrect,
		})

	case event.FinishedQuiz:
		quizAttempts, ok := p.quizAttempts[e.QuizID]
		if !ok {
			return fmt.Errorf("lastQuizAttempt %v not found", e.QuizID)
		}

		lastQuizAttempt := p.getLatestQuizAttempt(quizAttempts)

		err := lastQuizAttempt.checkFinishAttemptValidity()
		if err != nil {
			return err
		}

		lastQuizAttempt.completed = true

	default:
		panic(fmt.Sprintf("unknown event type %s", reflect.TypeOf(eventToApply)))
	}

	p.AppendEvent(eventToApply, isPersisted)

	return nil
}

func (p *Participant) getLatestQuizAttempt(quizAttempts []*quizAttempt) *quizAttempt {
	quizAttemptCount := len(quizAttempts)
	lastQuizAttempt := quizAttempts[quizAttemptCount-1]
	return lastQuizAttempt
}

func (p *Participant) ensureQuizNotStarted(id string) error {
	for _, quizAttempts := range p.quizAttempts {

		quizAttemptCount := len(quizAttempts)
		quiz := quizAttempts[quizAttemptCount-1]

		if quiz.QuizID == id && quiz.IsOngoing() {
			return fmt.Errorf("quiz '%s' already started and not finished", quiz.QuizID)
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

func (p *Participant) SelectQuizAnswer(quizID string, questionID string, answerID string, isCorrect bool) error {
	selectedAnswerEvent := event.SelectedAnswer{
		QuizID:     quizID,
		QuestionID: questionID,
		AnswerID:   answerID,
		IsCorrect:  isCorrect,
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

func (p *Participant) GetValidAttemptID(quizID string, attemptIDOrLatest string) (int, error) {
	var attemptIDNumber int
	quizAttemptCount := p.getQuizAttemptCountByQuizID(quizID)
	if attemptIDOrLatest == "latest" {
		attemptIDNumber = quizAttemptCount
	} else {
		attemptIDInt64, err := strconv.ParseInt(attemptIDOrLatest, 10, 0)
		if err != nil {
			return 0, err
		}
		attemptIDNumber = int(attemptIDInt64)
		if attemptIDNumber > quizAttemptCount {
			return 0, fmt.Errorf("attemptID %s request, but contains only %d attempts for quiz %s", attemptIDOrLatest, quizAttemptCount, quizID)
		}
	}

	return attemptIDNumber, nil
}

func (p *Participant) GetQuizAttemptCount(quizID string) int {
	return p.getQuizAttemptCountByQuizID(quizID)
}

func (p *Participant) getQuizAttemptCountByQuizID(quizID string) int {
	return len(p.quizAttempts[quizID])
}

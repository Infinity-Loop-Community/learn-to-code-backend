package projection

import (
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/domain/quiz/participant/event"
)

type QuizOverview struct {
	ActiveQuizzes   []string
	FinishedQuizzes []string
}

func NewQuizOverview(p participant.Participant) QuizOverview {

	qo := QuizOverview{
		ActiveQuizzes:   []string{},
		FinishedQuizzes: []string{},
	}

	activeQuizzesMap := map[string]string{}
	finishedQuizzesMap := map[string]string{}

	for _, generalEvent := range p.GetEvents() {

		switch e := generalEvent.(type) {

		case event.StartedQuiz:
			activeQuizzesMap[e.QuizID] = e.QuizID

		case event.FinishedQuiz:
			finishedQuizzesMap[e.QuizID] = e.QuizID
			delete(activeQuizzesMap, e.QuizID)
		}
	}

	for _, activeQuizID := range activeQuizzesMap {
		qo.ActiveQuizzes = append(qo.ActiveQuizzes, activeQuizID)
	}

	for _, finishedQuizID := range finishedQuizzesMap {
		qo.FinishedQuizzes = append(qo.FinishedQuizzes, finishedQuizID)
	}

	return qo
}

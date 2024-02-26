package projection

import (
	"fmt"
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/domain/quiz/participant/calculator"
	"learn-to-code/internal/domain/quiz/participant/event"
)

type QuizOverview struct {
	ActiveQuizzes   map[string][]QuizAttemptOverview
	FinishedQuizzes map[string][]QuizAttemptOverview
}

func (qo QuizOverview) GetFinishedQuizLatestAttempt(quizID string) (QuizAttemptOverview, error) {
	finishedQuiz, ok := qo.FinishedQuizzes[quizID]
	if !ok {
		return QuizAttemptOverview{}, fmt.Errorf("quiz %s does not exist", quizID)
	}

	if len(qo.FinishedQuizzes[quizID]) == 0 {
		return QuizAttemptOverview{}, fmt.Errorf("no quiz with id %s finished yet", quizID)
	}

	latestQuizAttempt := finishedQuiz[len(finishedQuiz)-1]

	return latestQuizAttempt, nil
}

func NewQuizOverview(p participant.Participant) (QuizOverview, error) {

	qo := QuizOverview{
		ActiveQuizzes:   map[string][]QuizAttemptOverview{},
		FinishedQuizzes: map[string][]QuizAttemptOverview{},
	}

	quizAttemptCounter := map[string]int{}
	quizFinishCounter := map[string]int{}

	quizResultCalculators := map[string]*calculator.QuizResult{}

	var activeQuizAttempts = map[string]*QuizAttemptOverview{}

	for _, generalEvent := range p.GetEvents() {

		switch e := generalEvent.(type) {

		case event.StartedQuiz:
			_, ok := quizAttemptCounter[e.QuizID]
			if !ok {
				quizAttemptCounter[e.QuizID] = 0
			}
			quizAttemptCounter[e.QuizID]++

			if activeQuizAttempts[e.QuizID] != nil {
				return QuizOverview{}, fmt.Errorf("invalid multiple active attempts for quiz %s", e.QuizID)
			}

			attemptAnswerCounterIndex := getQuizAttemptResultCalculatorKey(e.QuizID, quizAttemptCounter)
			_, ok = quizResultCalculators[attemptAnswerCounterIndex]
			if !ok {
				quizResultCalculators[attemptAnswerCounterIndex] = calculator.NewQuizResultCalculator()
			}

			activeQuizAttempts[e.QuizID] = &QuizAttemptOverview{
				QuizID:              e.QuizID,
				AttemptID:           quizAttemptCounter[e.QuizID],
				QuestionsWithAnswer: map[string]string{},
			}

		case event.SelectedAnswer:
			attemptAnswerCounterIndex := getQuizAttemptResultCalculatorKey(e.QuizID, quizAttemptCounter)

			if e.IsCorrect {
				quizResultCalculators[attemptAnswerCounterIndex].AddAnswer(e.QuestionID, true)
			} else {
				quizResultCalculators[attemptAnswerCounterIndex].AddAnswer(e.QuestionID, false)
			}

		case event.FinishedQuiz:
			_, ok := quizFinishCounter[e.QuizID]
			if !ok {
				quizFinishCounter[e.QuizID] = 0
			}
			quizFinishCounter[e.QuizID]++

			_, ok = qo.FinishedQuizzes[e.QuizID]
			if !ok {
				qo.FinishedQuizzes[e.QuizID] = []QuizAttemptOverview{}
			}

			calculatorKey := getQuizAttemptResultCalculatorKey(e.QuizID, quizAttemptCounter)
			activeQuizAttempts[e.QuizID].Pass = quizResultCalculators[calculatorKey].IsPass()

			activeQuizAttempts[e.QuizID].QuestionCorrectRatio = quizResultCalculators[calculatorKey].GetCorrectRatio()

			qo.FinishedQuizzes[e.QuizID] = append(qo.FinishedQuizzes[e.QuizID], *activeQuizAttempts[e.QuizID])
			delete(activeQuizAttempts, e.QuizID)
		}
	}

	for activeQuizID, activeQuizAttemptOverview := range activeQuizAttempts {
		_, ok := qo.ActiveQuizzes[activeQuizID]
		if !ok {
			qo.ActiveQuizzes[activeQuizID] = []QuizAttemptOverview{}
		}

		qo.ActiveQuizzes[activeQuizID] = append(qo.ActiveQuizzes[activeQuizID], QuizAttemptOverview{
			QuizID:              activeQuizID,
			AttemptID:           activeQuizAttemptOverview.AttemptID,
			QuestionsWithAnswer: activeQuizAttemptOverview.QuestionsWithAnswer,
			Pass:                false,
		})
	}

	return qo, nil
}

func getQuizAttemptResultCalculatorKey(quizID string, quizAttemptCounter map[string]int) string {
	return quizID + "-" + fmt.Sprintf("%d", quizAttemptCounter[quizID])
}

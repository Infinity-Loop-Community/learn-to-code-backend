package calculator

import "math"

type QuizResult struct {
	answerResults map[string]bool
}

const QuizPassThresold = 0.8

func NewQuizResultCalculator() *QuizResult {
	return &QuizResult{
		answerResults: map[string]bool{},
	}
}

func (qr *QuizResult) AddAnswer(questionID string, isCorrect bool) {
	qr.answerResults[questionID] = isCorrect
}

func (qr *QuizResult) GetCorrectRatio() float64 {
	if len(qr.answerResults) == 0 {
		return 0
	}

	correctAnswers := 0
	totalAnswers := len(qr.answerResults)

	for _, isCorrect := range qr.answerResults {
		if isCorrect {
			correctAnswers++
		}
	}

	return float64(correctAnswers) / float64(totalAnswers)
}

func (qr *QuizResult) IsPass() bool {
	if len(qr.answerResults) == 0 {
		return false
	}

	return qr.GetCorrectRatio() >= QuizPassThresold
}

func (qr *QuizResult) GetCorrectnessRatioComparedToOtherQuizResult(other *QuizResult) int {
	if len(qr.answerResults) == 0 || len(other.answerResults) == 0 {
		return 100
	}

	return int(math.Round((qr.GetCorrectRatio() - other.GetCorrectRatio()) * 100))
}

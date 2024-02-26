package quizattemptdetail

import (
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/domain/quiz/participant/calculator"
	"learn-to-code/internal/domain/quiz/participant/event"
	"math"
	"time"
)

type AttemptStatus string

const AttemptStatusOngoing = "ongoing"
const AttemptStatusFinished = "finished"

// taken from the last 1000 quizzes solved
const AverageTimePerQuestionMins = 2

type AttemptResult struct {
	Pass                                    bool
	QuestionCorrectRatio                    float64
	TimeTakenMins                           int
	ComparedToTimeAveragePercentage         int
	ComparedToCorrectRatioLastTryPercentage int
}

type QuizAttemptDetail struct {

	// AttemptID is created at runtime and start at 1, hence is human readable
	AttemptID int

	AttemptStatus AttemptStatus

	AttemptResult AttemptResult

	QuestionsWithAnswer map[string]string
}

func NewQuizAttemptDetail(p participant.Participant, quizID string, attemptID int) (QuizAttemptDetail, error) {

	qad := QuizAttemptDetail{
		QuestionsWithAnswer: map[string]string{},
	}

	quizCounter := 0

	startQuizTime := time.Time{}
	endQuizTime := time.Time{}

	prevQuizResultCalculator := calculator.NewQuizResultCalculator()
	quizResultCalculator := calculator.NewQuizResultCalculator()

	for _, generalEvent := range p.GetEvents() {

		switch e := generalEvent.(type) {

		case event.StartedQuiz:
			if e.QuizID == quizID {
				quizCounter++
				if (quizCounter) == attemptID {
					qad.AttemptID = quizCounter
					qad.AttemptStatus = AttemptStatusOngoing
					startQuizTime = e.CreatedAt
				}
			}

		case event.SelectedAnswer:
			if e.QuizID == quizID {

				if (quizCounter) == attemptID-1 {
					if e.IsCorrect {
						prevQuizResultCalculator.AddAnswer(e.QuestionID, true)
					} else {
						prevQuizResultCalculator.AddAnswer(e.QuestionID, false)
					}
				}

				if (quizCounter) == attemptID {
					qad.QuestionsWithAnswer[e.QuestionID] = e.AnswerID
					if e.IsCorrect {
						quizResultCalculator.AddAnswer(e.QuestionID, true)
					} else {
						quizResultCalculator.AddAnswer(e.QuestionID, false)
					}
				}
			}

		case event.FinishedQuiz:
			if e.QuizID == quizID {
				if (quizCounter) == attemptID {
					qad.AttemptStatus = AttemptStatusFinished
					endQuizTime = e.CreatedAt
				}
			}
		}

	}

	if (quizCounter) < attemptID {
		return QuizAttemptDetail{}, AttemptNotFoundError{
			AttemptID:    attemptID,
			AttemptCount: quizCounter,
			QuizID:       quizID,
		}
	}

	if attemptID < 1 {
		return QuizAttemptDetail{}, AttemptNotFoundError{
			AttemptID:    attemptID,
			AttemptCount: quizCounter,
			QuizID:       quizID,
		}
	}

	if qad.AttemptStatus == AttemptStatusFinished {

		comparedToCorrectRatioLastTryPercentage := quizResultCalculator.GetCorrectnessRatioComparedToOtherQuizResult(prevQuizResultCalculator)

		timeTakenMins := max(int(math.Round(endQuizTime.Sub(startQuizTime).Minutes())), 1)
		averageTimeMins := float64(len(qad.QuestionsWithAnswer) * AverageTimePerQuestionMins)
		comparedToTimeAveragePercentage := int(math.Round(((float64(timeTakenMins) / float64(averageTimeMins)) - 1) * 100))

		qad.AttemptResult = AttemptResult{
			Pass:                                    quizResultCalculator.IsPass(),
			QuestionCorrectRatio:                    quizResultCalculator.GetCorrectRatio(),
			TimeTakenMins:                           timeTakenMins,
			ComparedToTimeAveragePercentage:         comparedToTimeAveragePercentage,
			ComparedToCorrectRatioLastTryPercentage: comparedToCorrectRatioLastTryPercentage,
		}
	}

	return qad, nil
}

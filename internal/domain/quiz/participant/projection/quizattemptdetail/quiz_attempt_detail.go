package quizattemptdetail

import (
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/domain/quiz/participant/event"
	"math"
	"time"
)

type AttemptStatus string

const AttemptStatusOngoing = "ongoing"
const AttemptStatusFinished = "finished"
const QuizPassThresold = 0.8

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

	prevCorrectAnswerCounter := 0
	prevIncorrectAnswerCounter := 0

	correctAnswerCounter := 0
	incorrectAnswerCounter := 0

	startQuizTime := time.Time{}
	endQuizTime := time.Time{}

	for _, generalEvent := range p.Events {

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
						prevCorrectAnswerCounter++
					} else {
						prevIncorrectAnswerCounter++
					}
				}

				if (quizCounter) == attemptID {
					qad.QuestionsWithAnswer[e.QuestionID] = e.AnswerID
					if e.IsCorrect {
						correctAnswerCounter++
					} else {
						incorrectAnswerCounter++
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
		correctnessRatio := getCorrectnessRatio(correctAnswerCounter, incorrectAnswerCounter)
		prevCorrectnessRatio := getCorrectnessRatio(prevCorrectAnswerCounter, prevIncorrectAnswerCounter)
		comparedToCorrectRatioLastTryPercentage := int(math.Round((correctnessRatio - prevCorrectnessRatio) * 100))

		timeTakenMins := max(int(math.Round(endQuizTime.Sub(startQuizTime).Minutes())), 1)
		averageTimeMins := float64(len(qad.QuestionsWithAnswer) * AverageTimePerQuestionMins)
		comparedToTimeAveragePercentage := int(math.Round(((averageTimeMins / float64(timeTakenMins)) - 1) * 100))

		qad.AttemptResult = AttemptResult{
			Pass:                                    isPass(correctAnswerCounter, incorrectAnswerCounter),
			QuestionCorrectRatio:                    correctnessRatio,
			TimeTakenMins:                           timeTakenMins,
			ComparedToTimeAveragePercentage:         comparedToTimeAveragePercentage,
			ComparedToCorrectRatioLastTryPercentage: comparedToCorrectRatioLastTryPercentage,
		}
	}

	return qad, nil
}

func getCorrectnessRatio(correctAnswerCounter int, incorrectAnswerCounter int) float64 {
	if correctAnswerCounter+incorrectAnswerCounter == 0 {
		return 0
	}

	return (float64(correctAnswerCounter) / (float64(correctAnswerCounter) + float64(incorrectAnswerCounter)))
}

func isPass(correctAnswerCounter int, incorrectAnswerCounter int) bool {
	if correctAnswerCounter+incorrectAnswerCounter == 0 {
		return true
	}

	return getCorrectnessRatio(correctAnswerCounter, incorrectAnswerCounter) >= QuizPassThresold
}

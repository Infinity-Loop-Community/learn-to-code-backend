package projection

import (
	"fmt"
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/domain/quiz/participant/event"
)

type QuizAttemptDetail struct {
	QuestionsWithAnswer map[string]string
}

func NewQuizAttemptDetail(p participant.Participant, quizID string, attemptId int) (QuizAttemptDetail, error) {

	qad := QuizAttemptDetail{
		QuestionsWithAnswer: map[string]string{},
	}

	quizCounter := 0

	for _, generalEvent := range p.GetEvents() {

		switch e := generalEvent.(type) {

		case event.StartedQuiz:
			if e.QuizID == quizID {
				quizCounter++
			}

		case event.SelectedAnswer:
			if (quizCounter - 1) == attemptId {
				qad.QuestionsWithAnswer[e.QuestionID] = e.AnswerID
			}
		}
	}

	if (quizCounter - 1) < attemptId {
		return QuizAttemptDetail{}, fmt.Errorf("quiz attempt id %v does not exist, there are only %v attempts for quizId %v", attemptId, quizCounter, quizID)
	}

	return qad, nil
}

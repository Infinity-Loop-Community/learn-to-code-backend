package mapper

import (
	"learn-to-code/internal/domain/quiz/participant/projection"
	responseobject "learn-to-code/internal/interfaces/lambda/participant/quiz/responseobject"
)

type QuizAttemptDetailMapper struct {
}

func NewQuizAttemptDetailMapper() *QuizAttemptDetailMapper {
	return &QuizAttemptDetailMapper{}
}

func (cm *QuizAttemptDetailMapper) EntityToResponseObject(qad projection.QuizAttemptDetail) responseobject.QuizAttemptDetail {

	return responseobject.QuizAttemptDetail{
		QuestionsWithAnswer: qad.QuestionsWithAnswer,
	}
}

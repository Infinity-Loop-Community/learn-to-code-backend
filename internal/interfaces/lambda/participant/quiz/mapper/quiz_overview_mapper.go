package mapper

import (
	"learn-to-code/internal/domain/quiz/participant/projection"
	responseobject "learn-to-code/internal/interfaces/lambda/participant/quiz/responseobject"
)

type QuizOverviewMapper struct {
}

func NewQuizOverviewMapper() *QuizOverviewMapper {
	return &QuizOverviewMapper{}
}

func (cm *QuizOverviewMapper) EntityToResponseObject(p projection.QuizOverview) responseobject.QuizOverview {

	return responseobject.QuizOverview{
		ActiveQuizzes:   p.ActiveQuizzes,
		FinishedQuizzes: p.FinishedQuizzes,
	}
}

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
		ActiveQuizzes:   cm.toAttemptDetailResponseObjects(p.ActiveQuizzes),
		FinishedQuizzes: cm.toAttemptDetailResponseObjects(p.FinishedQuizzes),
	}
}

func (cm *QuizOverviewMapper) toAttemptDetailResponseObjects(attemptEntities map[string][]projection.QuizAttemptOverview) map[string][]responseobject.QuizAttemptOverview {
	attemptResponses := map[string][]responseobject.QuizAttemptOverview{}

	for quizID, attemptOverviewEntities := range attemptEntities {
		for _, attemptOverviewEntity := range attemptOverviewEntities {
			attemptResponses[quizID] = append(attemptResponses[quizID], responseobject.QuizAttemptOverview{
				AttemptID:           attemptOverviewEntity.AttemptID,
				QuizID:              attemptOverviewEntity.QuizID,
				QuestionsWithAnswer: attemptOverviewEntity.QuestionsWithAnswer,
			})
		}

	}

	return attemptResponses
}

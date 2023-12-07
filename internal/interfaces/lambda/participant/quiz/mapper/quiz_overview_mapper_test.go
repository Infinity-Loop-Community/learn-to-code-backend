package mapper_test

import (
	"learn-to-code/internal/domain/quiz/participant/projection"
	"learn-to-code/internal/infrastructure/inmemory"
	"learn-to-code/internal/infrastructure/testing/assertgo"
	"learn-to-code/internal/interfaces/lambda/participant/quiz/mapper"
	"strings"
	"testing"
)

func TestEntityToResponseObject(t *testing.T) {
	quizOverviewEntity := projection.QuizOverview{
		ActiveQuizzes: map[string][]projection.QuizAttemptOverview{
			inmemory.FirstQuestionID: {
				{
					QuizID:    inmemory.QuizID,
					AttemptID: 0,
					QuestionsWithAnswer: map[string]string{
						inmemory.FirstQuestionID: inmemory.FirstAnswerID,
					},
				},
			},
		},
		FinishedQuizzes: map[string][]projection.QuizAttemptOverview{
			inmemory.FirstQuestionID: {
				{
					QuizID:    inmemory.QuizID + "F",
					AttemptID: 1,
					QuestionsWithAnswer: map[string]string{
						inmemory.FirstQuestionID: inmemory.FirstAnswerID + "F",
					},
				},
			},
		},
	}

	qom := mapper.NewQuizOverviewMapper()
	resObj := qom.EntityToResponseObject(quizOverviewEntity)

	assertgo.NewAssertion(t, quizOverviewEntity).WithTypeKeyManipulation(strings.ToLower).IsEqualTo(resObj)
}

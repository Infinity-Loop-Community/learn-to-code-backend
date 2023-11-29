package mapper_test

import (
	"learn-to-code/internal/domain/quiz/course"
	"learn-to-code/internal/infrastructure/testing/assertgo"
	"learn-to-code/internal/interfaces/lambda/course/mapper"
	"strings"
	"testing"
)

func TestEntityToResponseObject(t *testing.T) {
	c := course.Course{
		ID:   "testID",
		Name: "testName",
		Steps: []course.Step{
			{
				ID:   "stepID",
				Name: "stepName",
				Quizzes: []course.StepQuiz{
					{
						ID: "quizID",
						Questions: []course.QuizQuestion{
							{
								Text:        "questionText",
								Difficulty:  "easy",
								Rating:      0,
								RatingCount: 0,
								Answers: []course.QuizAnswer{
									{
										Text:        "answerText",
										IsCorrect:   true,
										Description: "answerDescription",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	cm := mapper.NewCourseMapper()
	resObj := cm.EntityToResponseObject(c)

	assertgo.NewAssertion(t, c).WithTypeKeyManipulation(strings.ToLower).IsEqualTo(resObj)
}

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
								ID:          "16254bde-ac8c-409e-8a3e-f10805151a6b",
								Text:        "questionText",
								Difficulty:  "easy",
								Rating:      0,
								RatingCount: 0,
								Answers: []course.QuizAnswer{
									{
										ID:          "c362e4d9-f915-4480-bec0-488258e07186",
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

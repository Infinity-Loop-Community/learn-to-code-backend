package mapper

import (
	"learn-to-code/internal/domain/quiz/course"
	"learn-to-code/internal/interfaces/lambda/course/responseobject"
)

type CourseMapper struct {
}

func NewCourseMapper() *CourseMapper {
	return &CourseMapper{}
}

func (cm *CourseMapper) EntityToResponseObject(c course.Course) responseobject.Course {
	var responseSteps []responseobject.Step
	for _, s := range c.Steps {
		var responseQuizzes []responseobject.StepQuiz
		for _, q := range s.Quizzes {
			var responseQuestions []responseobject.QuizQuestion
			for _, qq := range q.Questions {
				var responseAnswers []responseobject.QuizAnswer
				for _, a := range qq.Answers {
					responseAnswers = append(responseAnswers, responseobject.QuizAnswer{
						ID:          a.ID,
						Text:        a.Text,
						IsCorrect:   a.IsCorrect,
						Description: a.Description,
					})
				}
				responseQuestions = append(responseQuestions, responseobject.QuizQuestion{
					ID:         qq.ID,
					Text:       qq.Text,
					Difficulty: qq.Difficulty,
					Answers:    responseAnswers,
				})
			}
			responseQuizzes = append(responseQuizzes, responseobject.StepQuiz{
				ID:        q.ID,
				Questions: responseQuestions,
			})
		}
		responseSteps = append(responseSteps, responseobject.Step{
			ID:      s.ID,
			Quizzes: responseQuizzes,
			Name:    s.Name,
		})
	}

	return responseobject.Course{
		ID:    c.ID,
		Steps: responseSteps,
		Name:  c.Name,
	}
}

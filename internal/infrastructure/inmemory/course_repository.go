package inmemory

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"learn-to-code/internal/domain/quiz/course"
	"learn-to-code/internal/interfaces/lambda/course/responseobject"
	"learn-to-code/static"
)

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{}
}

// CourseRepository contains hardcoded data for now to validate the requirements and access patterns
type CourseRepository struct {
}

const CourseIDFrontendDevelopment = "ed86d338-84a0-4486-a314-b99b17175875"

const CourseStepIDEssentialsOfTheWeb = "c7486278-a50c-4629-89b9-cc1c74d7a538"
const QuizIDEssentialsOfTheWeb = "fcf7890f-9c72-46d3-931e-34494307be37"
const FirstQuestionID = "14c20d31-c7e1-416d-9c8e-1f2040141f0b"
const FirstAnswerID = "06a1956e-b659-493f-9533-b27733ddd7fe"
const FirstCorrectAnswerID = "48a293ee-7f43-4e3d-85d1-4737e6385c7c"

const CourseStepIDJavaScriptBasics = "123f372e-e176-41c8-ba8e-9fc406c9ad1e"
const QuizIDJavaScriptBasics = "e3ce1f8b-bb40-4bdd-b31b-33cbef24d267"

func (q *CourseRepository) FindByID(courseID string) (course.Course, error) {

	quizEssentialsOfTheWeb, err := q.getQuiz(courseID, QuizIDEssentialsOfTheWeb)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	quizJavaScriptBasics, err := q.getQuiz(courseID, QuizIDJavaScriptBasics)
	if err != nil {
		return course.Course{}, course.ErrCourseNotFound
	}

	// Frontend Development
	if courseID == CourseIDFrontendDevelopment {
		return course.Course{
			ID:   CourseIDFrontendDevelopment,
			Name: "Frontend Development",
			Steps: []course.Step{
				{
					ID:      CourseStepIDEssentialsOfTheWeb,
					Name:    "The essentials of the Web",
					Quizzes: []course.StepQuiz{q.mapQuiz(quizEssentialsOfTheWeb)},
				},
				{
					ID:      CourseStepIDJavaScriptBasics,
					Name:    "JavaScript Basics",
					Quizzes: []course.StepQuiz{q.mapQuiz(quizJavaScriptBasics)},
				},
			},
		}, nil
	}

	return course.Course{}, course.ErrCourseNotFound
}

func (q *CourseRepository) getQuiz(courseID string, quizID string) (responseobject.StepQuiz, error) {
	file, err := q.readQuizFromFile(courseID, quizID)
	if err != nil {
		return responseobject.StepQuiz{}, err
	}

	stepQuiz, err := q.parseJSON(file)
	if err != nil {
		return responseobject.StepQuiz{}, err
	}

	return stepQuiz, nil
}

func (q *CourseRepository) parseJSON(file []byte) (responseobject.StepQuiz, error) {
	stepQuiz := responseobject.StepQuiz{}
	err := json.Unmarshal(file, &stepQuiz)
	if err != nil {
		return responseobject.StepQuiz{}, err
	}

	return stepQuiz, nil
}

func (q *CourseRepository) readQuizFromFile(requestedCourseID string, requestedQuizID string) ([]byte, error) {
	file, err := fs.ReadFile(static.StaticJSONFiles, fmt.Sprintf("json/course/%s/quiz/%s.json", requestedCourseID, requestedQuizID))
	if err != nil {
		return nil, fmt.Errorf("Unknown course %s or quiz %s, err: %v", requestedCourseID, requestedQuizID, err)
	}

	return file, nil
}

func (q *CourseRepository) mapQuiz(quiz responseobject.StepQuiz) course.StepQuiz {
	return course.StepQuiz{
		ID:        quiz.ID,
		Questions: mapQuestions(quiz.Questions),
	}
}

func mapQuestions(responseObjectQuestions []responseobject.QuizQuestion) []course.QuizQuestion {
	questions := []course.QuizQuestion{}

	for _, responseObjectQuestion := range responseObjectQuestions {
		questions = append(questions, course.QuizQuestion{
			ID:          responseObjectQuestion.ID,
			Text:        responseObjectQuestion.Text,
			Difficulty:  responseObjectQuestion.Difficulty,
			Answers:     mapAnswers(responseObjectQuestion.Answers),
			Rating:      float64(responseObjectQuestion.Rating),
			RatingCount: responseObjectQuestion.RatingCount,
		})
	}

	return questions
}

func mapAnswers(responseObjectAnswers []responseobject.QuizAnswer) []course.QuizAnswer {
	answers := []course.QuizAnswer{}

	for _, responseObjectAnswer := range responseObjectAnswers {
		answers = append(answers, course.QuizAnswer{
			ID:          responseObjectAnswer.ID,
			Text:        responseObjectAnswer.Text,
			IsCorrect:   responseObjectAnswer.IsCorrect,
			Description: responseObjectAnswer.Description,
		})
	}

	return answers
}

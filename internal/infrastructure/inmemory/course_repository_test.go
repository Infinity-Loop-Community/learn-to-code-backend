package inmemory

import (
	"testing"
)

func TestFindByID_ValidatesQuestionAndAnswerMapping(t *testing.T) {
	repo := NewCourseRepository()
	c, err := repo.FindByID(CourseIDFrontendDevelopment)

	if err != nil {
		t.Fatalf("FindByID returned an error: %v", err)
	}

	if len(c.Steps[0].Quizzes[0].Questions) == 0 {
		t.Errorf("No questions were parsed")
	}

	question := c.Steps[0].Quizzes[0].Questions[0]

	if question.ID == "" || question.Text == "" {
		t.Errorf("Question fields are not correctly parsed")
	}

	if len(question.Answers) == 0 {
		t.Errorf("No answers were parsed for question: %s", question.ID)
	}

	for _, answer := range question.Answers {
		if answer.ID == "" || answer.Text == "" {
			t.Errorf("Answer fields are not correctly parsed for question: %s", question.ID)
		}
	}
}

func TestFindByID_ReturnsErrorForUnkownCourse(t *testing.T) {
	repo := NewCourseRepository()
	_, err := repo.FindByID("unknown")

	if err == nil {
		t.Fatalf("FindByID returned no error for unknown course")
	}
}

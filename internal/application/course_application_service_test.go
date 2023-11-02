package application_test

import (
	"errors"
	"learn-to-code/internal/application"
	"learn-to-code/internal/domain/quiz/course"
	errUtils "learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/inmemory"
	"testing"
)

func TestCourseApplicationService_GetExistingQuiz(t *testing.T) {
	as := application.NewCourseApplicationService(
		inmemory.NewCourseRepository(),
	)

	course := errUtils.PanicIfError1(as.GetCourseByID(inmemory.CourseID))
	if course.ID != inmemory.CourseID {
		t.Fatalf("unexpected course id '%s', should be '%s'", course.ID, inmemory.CourseID)
	}
}

func TestCourseApplicationService_ErrorsForNotExistingId(t *testing.T) {
	as := application.NewCourseApplicationService(
		inmemory.NewCourseRepository(),
	)

	_, err := as.GetCourseByID("invalid-id")
	if !errors.Is(err, course.ErrCourseNotFound) {
		t.Fatalf("nil error or unexpected error for invalid course id '%v'", err)
	}
}

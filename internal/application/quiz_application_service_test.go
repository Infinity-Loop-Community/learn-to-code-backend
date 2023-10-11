package application_test

import (
	"learn-to-code/internal/application"
	errUtils "learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/go/util/uuid"
	"learn-to-code/internal/infrastructure/inmemory"
	"testing"
)

func TestQuizApplicationService_StartQuiz(t *testing.T) {
	as := application.NewQuizApplicationService(
		inmemory.NewParticipantRepository(),
	)

	userID := uuid.MustNewRandomAsString()

	startedQuizCount := errUtils.PanicIfError1(as.GetStartedQuizCount(userID))
	if startedQuizCount != 0 {
		t.Fatalf("new user, started quiz count not 0")
	}

	quizID := uuid.MustNewRandomAsString()
	errUtils.PanicIfError(as.StartQuiz(userID, quizID))

	startedQuizCount = errUtils.PanicIfError1(as.GetStartedQuizCount(userID))
	if startedQuizCount != 1 {
		t.Fatalf("after starting a quiz, count not 1")
	}
}

func TestQuizApplicationService_EndQuiz(t *testing.T) {
	as := application.NewQuizApplicationService(
		inmemory.NewParticipantRepository(),
	)

	userID := uuid.MustNewRandomAsString()

	finishedQuizCount := errUtils.PanicIfError1(as.GetFinishedQuizCount(userID))
	if finishedQuizCount != 0 {
		t.Fatalf("new user, finished quiz count not 0")
	}

	quizID := uuid.MustNewRandomAsString()
	_ = as.StartQuiz(userID, quizID)
	_ = as.FinishQuiz(userID, quizID)

	finishedQuizCount = errUtils.PanicIfError1(as.GetFinishedQuizCount(userID))
	if finishedQuizCount != 1 {
		t.Fatalf("after finishing a quiz, count not 1")
	}
}

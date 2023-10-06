package application_test

import (
	"hello-world/internal/application"
	errUtils "hello-world/internal/infrastructure/go/util/err"
	"hello-world/internal/infrastructure/go/util/uuid"
	"hello-world/internal/infrastructure/inmemory"
	"testing"
)

func TestQuizApplicationService_StartQuiz(t *testing.T) {
	as := application.NewQuizApplicationService(
		inmemory.NewParticipantRepository(),
	)

	userId := uuid.MustNewRandomAsString()

	startedQuizCount := errUtils.PanicIfError1(as.GetStartedQuizCount(userId))
	if startedQuizCount != 0 {
		t.Fatalf("new user, started quiz count not 0")
	}

	quizId := uuid.MustNewRandomAsString()
	errUtils.PanicIfError(as.StartQuiz(userId, quizId))

	startedQuizCount = errUtils.PanicIfError1(as.GetStartedQuizCount(userId))
	if startedQuizCount != 1 {
		t.Fatalf("after starting a quiz, count not 1")
	}
}

func TestQuizApplicationService_EndQuiz(t *testing.T) {
	as := application.NewQuizApplicationService(
		inmemory.NewParticipantRepository(),
	)

	userId := uuid.MustNewRandomAsString()

	finishedQuizCount := errUtils.PanicIfError1(as.GetFinishedQuizCount(userId))
	if finishedQuizCount != 0 {
		t.Fatalf("new user, finished quiz count not 0")
	}

	quizId := uuid.MustNewRandomAsString()
	_ = as.StartQuiz(userId, quizId)
	_ = as.FinishQuiz(userId, quizId)

	finishedQuizCount = errUtils.PanicIfError1(as.GetFinishedQuizCount(userId))
	if finishedQuizCount != 1 {
		t.Fatalf("after finishing a quiz, count not 1")
	}
}

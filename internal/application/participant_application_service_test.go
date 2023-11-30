package application_test

import (
	"learn-to-code/internal/application"
	"learn-to-code/internal/domain/command"
	errUtils "learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/go/util/uuid"
	"learn-to-code/internal/infrastructure/inmemory"
	"testing"
)

var commandFactory = command.NewCommandFactory()
var as = application.NewPartcipantApplicationService(
	inmemory.NewParticipantRepository(),
	command.NewParticipantCommandApplier(),
)

func TestQuizApplicationService_StartQuiz(t *testing.T) {
	userID := uuid.MustNewRandomAsString()

	startedQuizCount := errUtils.PanicIfError1(as.GetStartedQuizCount(userID))
	if startedQuizCount != 0 {
		t.Fatalf("new user, started quiz count not 0")
	}

	quizID := uuid.MustNewRandomAsString()
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(quizID), userID))

	startedQuizCount = errUtils.PanicIfError1(as.GetStartedQuizCount(userID))
	if startedQuizCount != 1 {
		t.Fatalf("after starting a quiz, count not 1")
	}
}

func TestQuizApplicationService_SelectAnswer(t *testing.T) {
	userID := uuid.MustNewRandomAsString()

	startedQuizCount := errUtils.PanicIfError1(as.GetStartedQuizCount(userID))
	if startedQuizCount != 0 {
		t.Fatalf("new user, started quiz count not 0")
	}

	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(inmemory.QuizID), userID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateSelectAnswerCommand(inmemory.QuizID, inmemory.FirstQuestionID, inmemory.FirstAnswerID), userID))
}

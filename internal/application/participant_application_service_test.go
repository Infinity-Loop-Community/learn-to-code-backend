package application_test

import (
	"context"
	"learn-to-code/internal/application"
	"learn-to-code/internal/domain/command"
	"learn-to-code/internal/domain/quiz/participant/event"
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/dynamodb"
	errUtils "learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/go/util/uuid"
	"learn-to-code/internal/infrastructure/inmemory"
	"learn-to-code/pkg/test/db"
	"os"
	"testing"
)

var commandFactory *command.Factory
var as *application.ParticipantApplicationService
var participantRepository *dynamodb.ParticipantRepository

func TestMain(m *testing.M) {
	dbStarter := db.NewDynamoStarter()

	participantRepository = dynamodb.NewDynamoDbParticipantRepository(context.Background(), config.Test, dbStarter.CreateDynamoDbClient(true))
	as = application.NewPartcipantApplicationService(
		participantRepository,
		command.NewParticipantCommandApplier(inmemory.NewCourseRepository()),
	)

	os.Exit(m.Run())
}

func TestQuizApplicationService_StartQuiz(t *testing.T) {
	userID := uuid.MustNewRandomAsString()

	startedQuizCount := errUtils.PanicIfError1(as.GetStartedQuizCount(userID))
	if startedQuizCount != 0 {
		t.Fatalf("new user, started quiz count not 0")
	}

	quizID := uuid.MustNewRandomAsString()
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(quizID, []string{inmemory.FirstQuestionID}), userID))

	startedQuizCount = errUtils.PanicIfError1(as.GetStartedQuizCount(userID))
	if startedQuizCount != 1 {
		t.Fatalf("after starting a quiz, count not 1")
	}
}

func TestQuizApplicationService_MapsEventsCorrectly(t *testing.T) {
	userID := uuid.MustNewRandomAsString()

	startedQuizCount := errUtils.PanicIfError1(as.GetStartedQuizCount(userID))
	if startedQuizCount != 0 {
		t.Fatalf("new user, started quiz count not 0")
	}

	quizID := uuid.MustNewRandomAsString()
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(quizID, []string{inmemory.FirstQuestionID}), userID))

	events := errUtils.PanicIfError1(participantRepository.FindEventsByID(userID))
	quizStartedEvent := events[1].(event.StartedQuiz)

	if quizStartedEvent.QuizID == "" {
		t.Fatalf("startedQuiz event quiz id is empty, event value is: %v", quizStartedEvent)
	}
}

func TestQuizApplicationService_SelectAnswer(t *testing.T) {
	userID := uuid.MustNewRandomAsString()

	startedQuizCount := errUtils.PanicIfError1(as.GetStartedQuizCount(userID))
	if startedQuizCount != 0 {
		t.Fatalf("new user, started quiz count not 0")
	}

	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(inmemory.QuizID, []string{inmemory.FirstQuestionID}), userID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateSelectAnswerCommand(inmemory.QuizID, inmemory.FirstQuestionID, inmemory.FirstAnswerID), userID))
}

func TestQuizApplicationService_FinishQuiz(t *testing.T) {
	userID := uuid.MustNewRandomAsString()

	startedQuizCount := errUtils.PanicIfError1(as.GetStartedQuizCount(userID))
	if startedQuizCount != 0 {
		t.Fatalf("new user, started quiz count not 0")
	}

	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(inmemory.QuizID, []string{inmemory.FirstQuestionID}), userID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateSelectAnswerCommand(inmemory.QuizID, inmemory.FirstQuestionID, inmemory.FirstAnswerID), userID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateFinishQuizCommand(inmemory.QuizID), userID))
}

func TestParticipantApplicationService_GetQuizAttemptDetail_NoQuizFinished_ReturnsEmpty(t *testing.T) {
	participantID := uuid.MustNewRandomAsString()

	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(inmemory.QuizID, []string{inmemory.FirstQuestionID}), participantID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateSelectAnswerCommand(inmemory.QuizID, inmemory.FirstQuestionID, inmemory.FirstAnswerID), participantID))

	_, err := as.GetLatestQuizAttemptDetail(participantID, inmemory.QuizID)

	if err == nil {
		t.Fatalf("does not return error for requesting the first attempt without having any quiz finished")
	}
}

func TestParticipantApplicationService_GetQuizAttemptDetail_ReturnsFirstLatestAttempt(t *testing.T) {
	participantID := uuid.MustNewRandomAsString()

	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(inmemory.QuizID, []string{inmemory.FirstQuestionID}), participantID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateSelectAnswerCommand(inmemory.QuizID, inmemory.FirstQuestionID, inmemory.FirstAnswerID), participantID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateFinishQuizCommand(inmemory.QuizID), participantID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(inmemory.QuizID, []string{inmemory.FirstQuestionID}), participantID))

	attemptDetail1 := errUtils.PanicIfError1(as.GetLatestQuizAttemptDetail(participantID, inmemory.QuizID))

	if attemptDetail1.AttemptID != 1 {
		t.Fatalf("first attempt id is not 1, %d instead", attemptDetail1.AttemptID)
	}
}

func TestParticipantApplicationService_GetQuizAttemptDetail_ReturnsLatest(t *testing.T) {
	participantID := uuid.MustNewRandomAsString()

	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(inmemory.QuizID, []string{inmemory.FirstQuestionID}), participantID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateSelectAnswerCommand(inmemory.QuizID, inmemory.FirstQuestionID, inmemory.FirstAnswerID), participantID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateFinishQuizCommand(inmemory.QuizID), participantID))

	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(inmemory.QuizID, []string{inmemory.FirstQuestionID}), participantID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateSelectAnswerCommand(inmemory.QuizID, inmemory.FirstQuestionID, inmemory.FirstAnswerID), participantID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateFinishQuizCommand(inmemory.QuizID), participantID))

	attemptDetail2 := errUtils.PanicIfError1(as.GetLatestQuizAttemptDetail(participantID, inmemory.QuizID))

	if len(attemptDetail2.QuestionsWithAnswer) == 0 {
		t.Fatalf("no question and answers for latest quiz attempt")
	}

	if attemptDetail2.AttemptID == 1 {
		t.Fatalf("no question and answers for latest quiz attempt")
	}
}

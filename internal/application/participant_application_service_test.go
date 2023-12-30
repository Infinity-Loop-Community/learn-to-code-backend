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
	"learn-to-code/internal/infrastructure/testing/db"
	"testing"
)

var commandFactory = command.NewCommandFactory()

func SetupApplicationService() (*application.ParticipantApplicationService, *dynamodb.ParticipantRepository, func()) {
	dynamoDbClient, clean := db.StartDynamoDB()

	participantRepository := dynamodb.NewDynamoDbParticipantRepository(context.Background(), config.Test, dynamoDbClient, dynamodb.NewEventPODeserializer())
	as := application.NewPartcipantApplicationService(
		participantRepository,
		command.NewParticipantCommandApplier(inmemory.NewCourseRepository()),
	)

	return as, participantRepository, clean
}

func TestQuizApplicationService_StartQuiz(t *testing.T) {
	as, _, clean := SetupApplicationService()
	defer clean()

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
	as, participantRepository, clean := SetupApplicationService()
	defer clean()

	userID := uuid.MustNewRandomAsString()

	startedQuizCount := errUtils.PanicIfError1(as.GetStartedQuizCount(userID))
	if startedQuizCount != 0 {
		t.Fatalf("new user, started quiz count not 0")
	}

	quizID := uuid.MustNewRandomAsString()
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(quizID, []string{inmemory.FirstQuestionID}), userID))

	events := errUtils.PanicIfError1(participantRepository.FindEventsByParticipantID(userID))
	quizStartedEvent := events[1].(event.StartedQuiz)

	if quizStartedEvent.QuizID == "" {
		t.Fatalf("startedQuiz event quiz id is empty, event value is: %v", quizStartedEvent)
	}
}

func TestQuizApplicationService_SelectAnswer(t *testing.T) {
	as, _, clean := SetupApplicationService()
	defer clean()

	userID := uuid.MustNewRandomAsString()

	startedQuizCount := errUtils.PanicIfError1(as.GetStartedQuizCount(userID))
	if startedQuizCount != 0 {
		t.Fatalf("new user, started quiz count not 0")
	}

	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(inmemory.QuizIDEssentialsOfTheWeb, []string{inmemory.FirstQuestionID}), userID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateSelectAnswerCommand(inmemory.QuizIDEssentialsOfTheWeb, inmemory.FirstQuestionID, inmemory.FirstAnswerID), userID))
}

func TestQuizApplicationService_FinishQuiz(t *testing.T) {
	as, _, clean := SetupApplicationService()
	defer clean()

	userID := uuid.MustNewRandomAsString()

	startedQuizCount := errUtils.PanicIfError1(as.GetStartedQuizCount(userID))
	if startedQuizCount != 0 {
		t.Fatalf("new user, started quiz count not 0")
	}

	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(inmemory.QuizIDEssentialsOfTheWeb, []string{inmemory.FirstQuestionID}), userID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateSelectAnswerCommand(inmemory.QuizIDEssentialsOfTheWeb, inmemory.FirstQuestionID, inmemory.FirstAnswerID), userID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateFinishQuizCommand(inmemory.QuizIDEssentialsOfTheWeb), userID))
}

func TestParticipantApplicationService_GetQuizAttemptDetail_NoQuizFinished_ReturnsEmpty(t *testing.T) {
	as, _, clean := SetupApplicationService()
	defer clean()

	participantID := uuid.MustNewRandomAsString()

	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(inmemory.QuizIDEssentialsOfTheWeb, []string{inmemory.FirstQuestionID}), participantID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateSelectAnswerCommand(inmemory.QuizIDEssentialsOfTheWeb, inmemory.FirstQuestionID, inmemory.FirstAnswerID), participantID))

	_, err := as.GetLatestQuizAttemptDetail(participantID, inmemory.QuizIDEssentialsOfTheWeb)

	if err == nil {
		t.Fatalf("does not return error for requesting the first attempt without having any quiz finished")
	}
}

func TestParticipantApplicationService_GetQuizAttemptDetail_ReturnsFirstLatestAttempt(t *testing.T) {
	as, _, clean := SetupApplicationService()
	defer clean()

	participantID := uuid.MustNewRandomAsString()

	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(inmemory.QuizIDEssentialsOfTheWeb, []string{inmemory.FirstQuestionID}), participantID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateSelectAnswerCommand(inmemory.QuizIDEssentialsOfTheWeb, inmemory.FirstQuestionID, inmemory.FirstAnswerID), participantID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateFinishQuizCommand(inmemory.QuizIDEssentialsOfTheWeb), participantID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(inmemory.QuizIDEssentialsOfTheWeb, []string{inmemory.FirstQuestionID}), participantID))

	attemptDetail1 := errUtils.PanicIfError1(as.GetLatestQuizAttemptDetail(participantID, inmemory.QuizIDEssentialsOfTheWeb))

	if attemptDetail1.AttemptID != 1 {
		t.Fatalf("first attempt id is not 1, %d instead", attemptDetail1.AttemptID)
	}
}

func TestParticipantApplicationService_GetQuizAttemptDetail_ReturnsLatest(t *testing.T) {
	as, _, clean := SetupApplicationService()
	defer clean()

	participantID := uuid.MustNewRandomAsString()

	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(inmemory.QuizIDEssentialsOfTheWeb, []string{inmemory.FirstQuestionID}), participantID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateSelectAnswerCommand(inmemory.QuizIDEssentialsOfTheWeb, inmemory.FirstQuestionID, inmemory.FirstAnswerID), participantID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateFinishQuizCommand(inmemory.QuizIDEssentialsOfTheWeb), participantID))

	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateStartQuizCommand(inmemory.QuizIDEssentialsOfTheWeb, []string{inmemory.FirstQuestionID}), participantID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateSelectAnswerCommand(inmemory.QuizIDEssentialsOfTheWeb, inmemory.FirstQuestionID, inmemory.FirstAnswerID), participantID))
	errUtils.PanicIfError(as.ProcessCommand(commandFactory.CreateFinishQuizCommand(inmemory.QuizIDEssentialsOfTheWeb), participantID))

	attemptDetail2 := errUtils.PanicIfError1(as.GetLatestQuizAttemptDetail(participantID, inmemory.QuizIDEssentialsOfTheWeb))

	if len(attemptDetail2.QuestionsWithAnswer) == 0 {
		t.Fatalf("no question and answers for latest quiz attempt")
	}

	if attemptDetail2.AttemptID == 1 {
		t.Fatalf("no question and answers for latest quiz attempt")
	}
}

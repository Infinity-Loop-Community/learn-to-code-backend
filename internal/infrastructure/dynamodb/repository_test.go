package dynamodb_test

import (
	"context"
	"errors"
	"learn-to-code/internal/domain/quiz/participant"
	dynamodb "learn-to-code/internal/infrastructure/dynamodb"
	errUtils "learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/go/util/uuid"
	"learn-to-code/pkg/test/db"
	"os"
	"testing"

	dynamodbsdk "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var dynamoDbClient *dynamodbsdk.Client

func TestMain(m *testing.M) {
	dbStarter := db.NewDynamoStarter()
	dynamoDbClient = dbStarter.Start()

	defer dbStarter.Terminate()

	// now we just need to tell go-test that we can run the tests
	os.Exit(m.Run())
}

func TestRepository_FindById_ReturnsNotFoundError(t *testing.T) {
	repo := getRepository()

	_, err := repo.FindByID("does not exist")

	if !errors.Is(err, participant.ErrParticipantNotFound) {
		t.Fatalf("is not ErrParticipantNotFound: %s", err)
	}
}

func TestRepository_FindById_HandleSingleUser(t *testing.T) {
	repo := getRepository()

	p := errUtils.PanicIfError1(participant.New())

	quizID := uuid.MustNewRandomAsString()
	errUtils.PanicIfError(
		p.StartQuiz(quizID),
	)

	errUtils.PanicIfError(
		repo.AppendEvents(p.GetID(), p.GetNewEventsAndUpdatePersistedVersion()),
	)

	p, err := repo.FindByID(p.GetID())

	if err != nil {
		t.Fatalf("could not fetch the participant due to an error: %s", err)
	}

	errUtils.PanicIfError(p.FinishQuiz(quizID))
	errUtils.PanicIfError(
		repo.AppendEvents(p.GetID(), p.GetNewEventsAndUpdatePersistedVersion()),
	)

	_, err = repo.FindByID(p.GetID())
	if err != nil {
		t.Fatalf("error while getting a participant with finished quiz: %s", err)
	}
}

func getRepository() participant.Repository {
	repo := dynamodb.NewDynamoDbParticipantRepository(context.Background(), "test", dynamoDbClient)

	return repo
}

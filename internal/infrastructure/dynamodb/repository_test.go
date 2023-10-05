package dynamodb_test

import (
	"errors"
	"hello-world/internal/domain/quiz/participant"
	dynamodb "hello-world/internal/infrastructure/dynamodb"
	errUtils "hello-world/internal/infrastructure/go/util/err"
	"hello-world/internal/infrastructure/go/util/uuid"
	"hello-world/pkg/test/db"
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

	_, err := repo.FindById("does not exist")

	if !errors.Is(err, participant.ErrNotFound) {
		t.Fatalf("is not ErrNotFound: %s", err)
	}
}

func TestRepository_FindById_HandleSingleUser(t *testing.T) {
	repo := getRepository()

	p := errUtils.PanicIfError1(participant.New())

	quizId := uuid.MustNewRandomAsString()
	errUtils.PanicIfError(
		p.StartQuiz(quizId),
	)

	errUtils.PanicIfError(
		repo.AppendEvents(p.GetId(), p.GetNewEventsAndUpdatePersistedVersion()),
	)

	p, err := repo.FindById(p.GetId())

	if err != nil {
		t.Fatalf("could not fetch the participant due to an error: %s", err)
	}

	errUtils.PanicIfError(p.FinishQuiz(quizId))
	errUtils.PanicIfError(
		repo.AppendEvents(p.GetId(), p.GetNewEventsAndUpdatePersistedVersion()),
	)

	_, err = repo.FindById(p.GetId())
	if err != nil {
		t.Fatalf("error while getting a participant with finished quiz: %s", err)
	}
}

func getRepository() participant.Repository {
	var repo participant.Repository

	repo = dynamodb.NewDynamoDbParticipantRepository(dynamoDbClient)

	return repo
}

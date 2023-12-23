package dynamodb_test

import (
	"context"
	"learn-to-code/internal/domain/quiz/participant"
	dynamodb "learn-to-code/internal/infrastructure/dynamodb"
	errUtils "learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/go/util/uuid"
	"learn-to-code/internal/infrastructure/inmemory"
	"learn-to-code/pkg/test/db"
	"testing"
)

func TestRepository_FindById_ReturnsNewUser(t *testing.T) {
	repo := getRepository()

	p, err := repo.FindOrCreateByID("does not exist")

	if err != nil {
		t.Fatalf("error finding a user who does not exist: %s", err)
	}

	if p.GetID() == "" {
		t.Fatalf("new user id is nil")
	}
}

func TestRepository_FindById_HandleSingleUser(t *testing.T) {
	repo := getRepository()

	p := errUtils.PanicIfError1(participant.New())

	quizID := uuid.MustNewRandomAsString()
	errUtils.PanicIfError(
		p.StartQuiz(quizID, nil),
	)

	errUtils.PanicIfError(
		repo.StoreEvents(p.GetID(), p.GetNewEventsAndUpdatePersistedVersion()),
	)

	p, err := repo.FindOrCreateByID(p.GetID())

	if err != nil {
		t.Fatalf("could not fetch the participant due to an error: %s", err)
	}

	errUtils.PanicIfError(p.SelectQuizAnswer(quizID, inmemory.FirstQuestionID, inmemory.FirstAnswerID, true))

	errUtils.PanicIfError(p.FinishQuiz(quizID))
	errUtils.PanicIfError(
		repo.StoreEvents(p.GetID(), p.GetNewEventsAndUpdatePersistedVersion()),
	)

	_, err = repo.FindOrCreateByID(p.GetID())
	if err != nil {
		t.Fatalf("error while getting a participant with finished quiz: %s", err)
	}
}

func TestParticipantRepository_StoreEventsWithPayload(t *testing.T) {
	repo := getRepository()

	p := errUtils.PanicIfError1(participant.New())

	quizID := uuid.MustNewRandomAsString()
	errUtils.PanicIfError(
		p.StartQuiz(quizID, nil),
	)

	errUtils.PanicIfError(
		p.SelectQuizAnswer(quizID, inmemory.FirstQuestionID, inmemory.FirstAnswerID, true),
	)

	errUtils.PanicIfError(
		repo.StoreEvents(p.GetID(), p.GetNewEventsAndUpdatePersistedVersion()),
	)

	events := errUtils.PanicIfError1(repo.FindEventsByID(p.GetID()))

	for _, event := range events {
		if event.GetAggregateID() == "" {
			t.Fatalf("aggregateID is empty")
		}

		if event.GetCreatedAt().String() == "" {
			t.Fatalf("createdAt is empty")
		}
	}
}

func getRepository() participant.Repository {
	dbStarter := db.NewDynamoStarter()

	repo := dynamodb.NewDynamoDbParticipantRepository(context.Background(), "test", dbStarter.CreateDynamoDbClient(true))

	return repo
}

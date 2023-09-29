package participant_test

import (
	"errors"
	dParticipant "hello-world/internal/domain/quiz/participant"
	errUtils "hello-world/internal/infrastructure/go/util/err"
	"hello-world/internal/infrastructure/go/util/uuid"
	iParticipant "hello-world/internal/infrastructure/inmemory"
	"testing"
)

func TestRepository_FindById_ReturnsNotFoundError(t *testing.T) {
	var repo dParticipant.Repository
	repo = iParticipant.NewParticipantRepository()

	_, err := repo.FindById("does not exist")

	if !errors.Is(err, dParticipant.ErrNotFound) {
		t.Fatalf("is not ErrNotFound")
	}
}

func TestRepository_FindById_HandleSingleUser(t *testing.T) {
	var repo dParticipant.Repository
	repo = iParticipant.NewParticipantRepository()

	p := dParticipant.New()
	quizStartedEvent := errUtils.PanicIfError1(
		p.StartQuiz(uuid.MustNewRandomAsString()),
	)

	errUtils.PanicIfError(
		repo.AppendEvent(p.GetId(), quizStartedEvent),
	)

	_, err := repo.FindById(p.GetId())

	if errors.Is(err, dParticipant.ErrNotFound) {
		t.Fatalf("is ErrNotFound, participant should exist")
	}
}

func TestName(t *testing.T) {

}

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

	p := errUtils.PanicIfError1(dParticipant.New())

	quizId := uuid.MustNewRandomAsString()
	quizStartedEvent := errUtils.PanicIfError1(
		p.StartQuiz(quizId),
	)

	errUtils.PanicIfError(
		repo.AppendEvent(p.GetId(), quizStartedEvent),
	)

	p, err := repo.FindById(p.GetId())

	if err != nil {
		t.Fatalf("could not fetch the participant due to an error: %s", err)
	}

	finishedQuizEvent := errUtils.PanicIfError1(p.FinishQuiz(quizId))
	errUtils.PanicIfError(
		repo.AppendEvent(p.GetId(), finishedQuizEvent),
	)

	_, err = repo.FindById(p.GetId())
	if err != nil {
		t.Fatalf("error while getting a participant with finished quiz: %s", err)
	}
}

func TestName(t *testing.T) {

}

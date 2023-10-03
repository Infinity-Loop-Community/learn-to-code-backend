package participant_test

import (
	"github.com/google/uuid"
	"hello-world/internal/domain/quiz/participant"
	"hello-world/internal/infrastructure/go/util/err"
	"strings"
	"testing"
)

func TestParticipant_GetId_CorrectSize(t *testing.T) {
	p := err.PanicIfError1(participant.New())

	pLength := len(p.GetId())
	uuidLength := len(newUuid())

	if pLength != uuidLength {
		t.Fatalf("Id length not equal to an uuid %d != %d", pLength, uuidLength)
	}
}

func TestParticipant_GetId_Unique(t *testing.T) {
	p1 := err.PanicIfError1(participant.New())
	p2 := err.PanicIfError1(participant.New())

	if p1.GetId() == p2.GetId() {
		t.Fatalf("Id not unique")
	}
}

func TestParticipant_GetJoinedQuizzesInitiallyZero(t *testing.T) {
	p := err.PanicIfError1(participant.New())

	if p.GetStartedQuizCount() != 0 {
		t.Fatalf("joined quiz count initially not 0")
	}
}

func TestParticipant_WhenJoiningAQuiz_HasJoinedQuizzes(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	_, _ = p.StartQuiz(quizId)

	if p.GetStartedQuizCount() != 1 {
		t.Fatalf("after joining a quiz count is not 1")
	}
}

func TestParticipant_JoinQuiz_ReturnsEvent(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	event, _ := p.StartQuiz(quizId)

	if event.Id != quizId {
		t.Fatalf("invalid event Id '%s' != '%s", event.Id, quizId)
	}
}

func TestParticipant_JoinQuiz_FailsIfJoinedTwice(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	_, err1 := p.StartQuiz(quizId)
	if err1 != nil {
		t.Fatalf("initial start quiz should not fail")
	}

	_, err2 := p.StartQuiz(quizId)
	if err2 == nil {
		t.Fatalf("starting a quiz twice without finishing should fail")
	}
}

func TestParticipant_JoinQuiz_JoiningTheQuizAgainSucceedsAfterFinishAgain(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	_, err1 := p.StartQuiz(quizId)
	if err1 != nil {
		t.Fatalf("initial start quiz should not fail")
	}

	_, _ = p.FinishQuiz(quizId)

	_, err2 := p.StartQuiz(quizId)
	if err2 != nil {
		t.Fatalf("initial start quiz should not fail")
	}

	if p.GetStartedQuizCount() != 2 {
		t.Fatalf("start quiz count not 2 although started twice")
	}

	_, finishQuizErr := p.FinishQuiz(quizId)
	if finishQuizErr != nil {
		t.Fatalf("finish a quiz before starting does not fail")
	}
}

func TestParticipant_FinishQuiz_FailsIfNotStarted(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	_, finishQuizErr := p.FinishQuiz(quizId)

	if finishQuizErr == nil || strings.Contains(finishQuizErr.Error(), "no started") {
		t.Fatalf("finish a quiz before starting does not fail")
	}
}

func TestParticipant_FinishQuiz(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	_, _ = p.StartQuiz(quizId)
	_, finishQuizErr := p.FinishQuiz(quizId)

	if finishQuizErr != nil {
		t.Fatalf("finish quiz failed")
	}
}

func TestParticipant_FinishQuiz_failsIfFinishedTwice(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	_, _ = p.StartQuiz(quizId)
	_, finishQuizErr1 := p.FinishQuiz(quizId)
	_, finishQuizErr2 := p.FinishQuiz(quizId)

	if finishQuizErr1 != nil {
		t.Fatalf("first finish quiz failed")
	}

	if finishQuizErr2 == nil {
		t.Fatalf("quiz finished twice but did not error")
	}
}

func TestParticipant_FinishQuiz_eventQuizIdMatches(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	_, _ = p.StartQuiz(quizId)
	event, _ := p.FinishQuiz(quizId)

	if event.Id != quizId {
		t.Fatalf("finish event quizId does not match quiz Id '%s' != '%s'", event.Id, quizId)
	}
}

func TestParticipant_GetFinishedQuizCount(t *testing.T) {
	p := err.PanicIfError1(participant.New())

	if p.GetFinishedQuizCount() != 0 {
		t.Fatalf("finished quiz count must be 0 initially")
	}
}

func TestParticipant_FinishQuiz_GetFinishedQuizCount(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	_, _ = p.StartQuiz(quizId)
	_, finishQuizErr := p.FinishQuiz(quizId)

	if finishQuizErr != nil {
		t.Fatalf("finish quiz error")
	}

	if p.GetFinishedQuizCount() != 1 {
		t.Fatalf("finished quiz count must be 1 after finishing 1 quiz")
	}
}

func newUuid() string {
	return err.PanicIfError1(uuid.NewRandom()).String()
}

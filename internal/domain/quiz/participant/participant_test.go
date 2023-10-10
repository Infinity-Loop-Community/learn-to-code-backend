package participant_test

import (
	"encoding/json"
	"github.com/google/uuid"
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/infrastructure/go/util/err"
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

	_ = p.StartQuiz(quizId)

	if p.GetStartedQuizCount() != 1 {
		t.Fatalf("after joining a quiz count is not 1")
	}
}

func TestParticipant_JoinQuiz_ReturnsEvent(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	_ = p.StartQuiz(quizId)

	newEvents := p.GetNewEventsAndUpdatePersistedVersion()

	if newEvents[len(newEvents)-1].GetId() != quizId {
		t.Fatalf("invalid event Id '%s' != '%s", newEvents[len(newEvents)-1].GetId(), quizId)
	}
}

func TestParticipant_JoinQuiz_FailsIfJoinedTwice(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	err1 := p.StartQuiz(quizId)
	if err1 != nil {
		t.Fatalf("initial start quiz should not fail")
	}

	err2 := p.StartQuiz(quizId)
	if err2 == nil {
		t.Fatalf("starting a quiz twice without finishing should fail")
	}
}

func TestParticipant_JoinQuiz_JoiningTheQuizAgainSucceedsAfterFinishAgain(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	err1 := p.StartQuiz(quizId)
	if err1 != nil {
		t.Fatalf("initial start quiz should not fail")
	}

	_ = p.FinishQuiz(quizId)

	err2 := p.StartQuiz(quizId)
	if err2 != nil {
		t.Fatalf("initial start quiz should not fail")
	}

	if p.GetStartedQuizCount() != 2 {
		t.Fatalf("start quiz count not 2 although started twice")
	}

	finishQuizErr := p.FinishQuiz(quizId)
	if finishQuizErr != nil {
		t.Fatalf("finish a quiz before starting does not fail")
	}
}

func TestParticipant_FinishQuiz_FailsIfNotStarted(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	finishQuizErr := p.FinishQuiz(quizId)

	if finishQuizErr == nil || strings.Contains(finishQuizErr.Error(), "no started") {
		t.Fatalf("finish a quiz before starting does not fail")
	}
}

func TestParticipant_FinishQuiz(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	_ = p.StartQuiz(quizId)
	finishQuizErr := p.FinishQuiz(quizId)

	if finishQuizErr != nil {
		t.Fatalf("finish quiz failed")
	}
}

func TestParticipant_FinishQuiz_failsIfFinishedTwice(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	_ = p.StartQuiz(quizId)
	finishQuizErr1 := p.FinishQuiz(quizId)
	finishQuizErr2 := p.FinishQuiz(quizId)

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

	err.PanicIfError(p.StartQuiz(quizId))
	err.PanicIfError(p.FinishQuiz(quizId))

	newEvents := p.GetNewEventsAndUpdatePersistedVersion()
	lastEvent := newEvents[len(newEvents)-1]

	if lastEvent.GetId() != quizId {
		t.Fatalf("finish event quizId does not match quiz Id '%s' != '%s'", lastEvent.GetId(), quizId)
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

	_ = p.StartQuiz(quizId)
	finishQuizErr := p.FinishQuiz(quizId)

	if finishQuizErr != nil {
		t.Fatalf("finish quiz error")
	}

	if p.GetFinishedQuizCount() != 1 {
		t.Fatalf("finished quiz count must be 1 after finishing 1 quiz")
	}
}

func TestParticipant_Events_applyAndRestoresWithSameVersion(t *testing.T) {

	p1 := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	err.PanicIfError(p1.StartQuiz(quizId))
	err.PanicIfError(p1.FinishQuiz(quizId))

	participantEvents := p1.Events

	p2 := err.PanicIfError1(participant.NewFromEvents(participantEvents))

	if p1.CurrentVersion != p2.CurrentVersion {
		t.Fatalf("original participant's version is different to the restored version: %d != %d", p1.CurrentVersion, p2.CurrentVersion)
	}

}

func TestParticipant_Events_createsSameEventsAfterApplyAndRestore(t *testing.T) {

	p1 := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	err.PanicIfError(p1.StartQuiz(quizId))
	err.PanicIfError(p1.FinishQuiz(quizId))

	participantEvents := p1.Events

	p2 := err.PanicIfError1(participant.NewFromEvents(participantEvents))

	quiz2Id := newUuid()
	err.PanicIfError(p1.StartQuiz(quiz2Id))
	err.PanicIfError(p2.StartQuiz(quiz2Id))

	p1NewEvents := p1.GetNewEventsAndUpdatePersistedVersion()
	p1LastEvent := p1NewEvents[len(p1NewEvents)-1]

	p2NewEvents := p2.GetNewEventsAndUpdatePersistedVersion()
	p2LastEvent := p2NewEvents[len(p2NewEvents)-1]

	if p1LastEvent.GetVersion() != p2LastEvent.GetVersion() || p1LastEvent.GetId() != p2LastEvent.GetId() {
		t.Fatalf("original participant's new event is different to the restored version: %v != %v", p1LastEvent, p2LastEvent)
	}
}

func TestParticipant_Events_applyAndRestoresWithSameEvents(t *testing.T) {

	p1 := err.PanicIfError1(participant.New())
	quizId := err.PanicIfError1(uuid.NewRandom()).String()

	err.PanicIfError(p1.StartQuiz(quizId))
	err.PanicIfError(p1.FinishQuiz(quizId))

	participantEvents := p1.Events

	p2 := err.PanicIfError1(participant.NewFromEvents(participantEvents))

	p1EventsAsJson := string(err.PanicIfError1(json.Marshal(p1.Events)))
	p2EventsAsJson := string(err.PanicIfError1(json.Marshal(p2.Events)))
	if p1EventsAsJson != p2EventsAsJson {
		t.Fatalf("original participant's events are different to the restored version: %s != %s", p1EventsAsJson, p2EventsAsJson)
	}
}

func newUuid() string {
	return err.PanicIfError1(uuid.NewRandom()).String()
}

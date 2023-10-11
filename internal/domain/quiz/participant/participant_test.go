package participant_test

import (
	"encoding/json"
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/infrastructure/go/util/err"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestParticipant_GetId_CorrectSize(t *testing.T) {
	p := err.PanicIfError1(participant.New())

	pLength := len(p.GetID())
	uuidLength := len(newUUID())

	if pLength != uuidLength {
		t.Fatalf("ID length not equal to an uuid %d != %d", pLength, uuidLength)
	}
}

func TestParticipant_GetId_Unique(t *testing.T) {
	p1 := err.PanicIfError1(participant.New())
	p2 := err.PanicIfError1(participant.New())

	if p1.GetID() == p2.GetID() {
		t.Fatalf("ID not unique")
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
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	_ = p.StartQuiz(quizID)

	if p.GetStartedQuizCount() != 1 {
		t.Fatalf("after joining a quiz count is not 1")
	}
}

func TestParticipant_JoinQuiz_ReturnsEvent(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	_ = p.StartQuiz(quizID)

	newEvents := p.GetNewEventsAndUpdatePersistedVersion()

	if newEvents[len(newEvents)-1].GetID() != quizID {
		t.Fatalf("invalid event ID '%s' != '%s", newEvents[len(newEvents)-1].GetID(), quizID)
	}
}

func TestParticipant_JoinQuiz_FailsIfJoinedTwice(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	err1 := p.StartQuiz(quizID)
	if err1 != nil {
		t.Fatalf("initial start quiz should not fail")
	}

	err2 := p.StartQuiz(quizID)
	if err2 == nil {
		t.Fatalf("starting a quiz twice without finishing should fail")
	}
}

func TestParticipant_JoinQuiz_JoiningTheQuizAgainSucceedsAfterFinishAgain(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	err1 := p.StartQuiz(quizID)
	if err1 != nil {
		t.Fatalf("initial start quiz should not fail")
	}

	_ = p.FinishQuiz(quizID)

	err2 := p.StartQuiz(quizID)
	if err2 != nil {
		t.Fatalf("initial start quiz should not fail")
	}

	if p.GetStartedQuizCount() != 2 {
		t.Fatalf("start quiz count not 2 although started twice")
	}

	finishQuizErr := p.FinishQuiz(quizID)
	if finishQuizErr != nil {
		t.Fatalf("finish a quiz before starting does not fail")
	}
}

func TestParticipant_FinishQuiz_FailsIfNotStarted(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	finishQuizErr := p.FinishQuiz(quizID)

	if finishQuizErr == nil || strings.Contains(finishQuizErr.Error(), "no started") {
		t.Fatalf("finish a quiz before starting does not fail")
	}
}

func TestParticipant_FinishQuiz(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	_ = p.StartQuiz(quizID)
	finishQuizErr := p.FinishQuiz(quizID)

	if finishQuizErr != nil {
		t.Fatalf("finish quiz failed")
	}
}

func TestParticipant_FinishQuiz_failsIfFinishedTwice(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	_ = p.StartQuiz(quizID)
	finishQuizErr1 := p.FinishQuiz(quizID)
	finishQuizErr2 := p.FinishQuiz(quizID)

	if finishQuizErr1 != nil {
		t.Fatalf("first finish quiz failed")
	}

	if finishQuizErr2 == nil {
		t.Fatalf("quiz finished twice but did not error")
	}
}

func TestParticipant_FinishQuiz_eventQuizIdMatches(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	err.PanicIfError(p.StartQuiz(quizID))
	err.PanicIfError(p.FinishQuiz(quizID))

	newEvents := p.GetNewEventsAndUpdatePersistedVersion()
	lastEvent := newEvents[len(newEvents)-1]

	if lastEvent.GetID() != quizID {
		t.Fatalf("finish event quizId does not match quiz ID '%s' != '%s'", lastEvent.GetID(), quizID)
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
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	_ = p.StartQuiz(quizID)
	finishQuizErr := p.FinishQuiz(quizID)

	if finishQuizErr != nil {
		t.Fatalf("finish quiz error")
	}

	if p.GetFinishedQuizCount() != 1 {
		t.Fatalf("finished quiz count must be 1 after finishing 1 quiz")
	}
}

func TestParticipant_Events_applyAndRestoresWithSameVersion(t *testing.T) {

	p1 := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	err.PanicIfError(p1.StartQuiz(quizID))
	err.PanicIfError(p1.FinishQuiz(quizID))

	participantEvents := p1.Events

	p2 := err.PanicIfError1(participant.NewFromEvents(participantEvents))

	if p1.CurrentVersion != p2.CurrentVersion {
		t.Fatalf("original participant's version is different to the restored version: %d != %d", p1.CurrentVersion, p2.CurrentVersion)
	}

}

func TestParticipant_Events_createsSameEventsAfterApplyAndRestore(t *testing.T) {

	p1 := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	err.PanicIfError(p1.StartQuiz(quizID))
	err.PanicIfError(p1.FinishQuiz(quizID))

	participantEvents := p1.Events

	p2 := err.PanicIfError1(participant.NewFromEvents(participantEvents))

	quiz2Id := newUUID()
	err.PanicIfError(p1.StartQuiz(quiz2Id))
	err.PanicIfError(p2.StartQuiz(quiz2Id))

	p1NewEvents := p1.GetNewEventsAndUpdatePersistedVersion()
	p1LastEvent := p1NewEvents[len(p1NewEvents)-1]

	p2NewEvents := p2.GetNewEventsAndUpdatePersistedVersion()
	p2LastEvent := p2NewEvents[len(p2NewEvents)-1]

	if p1LastEvent.GetVersion() != p2LastEvent.GetVersion() || p1LastEvent.GetID() != p2LastEvent.GetID() {
		t.Fatalf("original participant's new event is different to the restored version: %v != %v", p1LastEvent, p2LastEvent)
	}
}

func TestParticipant_Events_applyAndRestoresWithSameEvents(t *testing.T) {

	p1 := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	err.PanicIfError(p1.StartQuiz(quizID))
	err.PanicIfError(p1.FinishQuiz(quizID))

	participantEvents := p1.Events

	p2 := err.PanicIfError1(participant.NewFromEvents(participantEvents))

	p1EventsAsJSON := string(err.PanicIfError1(json.Marshal(p1.Events)))
	p2EventsAsJSON := string(err.PanicIfError1(json.Marshal(p2.Events)))
	if p1EventsAsJSON != p2EventsAsJSON {
		t.Fatalf("original participant's events are different to the restored version: %s != %s", p1EventsAsJSON, p2EventsAsJSON)
	}
}

func newUUID() string {
	return err.PanicIfError1(uuid.NewRandom()).String()
}

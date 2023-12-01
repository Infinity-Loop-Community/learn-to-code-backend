package participant_test

import (
	"encoding/json"
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/domain/quiz/participant/event"
	"learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/inmemory"
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

	_ = p.StartQuiz(quizID, nil)

	if p.GetStartedQuizCount() != 1 {
		t.Fatalf("after joining a quiz count is not 1")
	}
}

func TestParticipant_JoinQuiz_ReturnsEvent(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	_ = p.StartQuiz(quizID, nil)

	newEvents := p.GetNewEventsAndUpdatePersistedVersion()

	startedQuizEvent := newEvents[len(newEvents)-1].(event.StartedQuiz)
	if startedQuizEvent.QuizID != quizID {
		t.Fatalf("invalid event ID '%s' != '%s", startedQuizEvent.QuizID, quizID)
	}
}

func TestParticipant_JoinQuiz_FailsIfJoinedTwice(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	err1 := p.StartQuiz(quizID, nil)
	if err1 != nil {
		t.Fatalf("initial start quiz should not fail")
	}

	err2 := p.StartQuiz(quizID, nil)
	if err2 == nil {
		t.Fatalf("starting a quiz twice without finishing should fail")
	}
}

func TestParticipant_JoinQuiz_JoiningTheQuizAgainSucceedsAfterFinishAgain(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	err1 := p.StartQuiz(quizID, nil)
	if err1 != nil {
		t.Fatalf("initial start quiz should not fail")
	}

	_ = p.FinishQuiz(quizID)

	err2 := p.StartQuiz(quizID, nil)
	if err2 != nil {
		t.Fatalf("initial start quiz should not fail")
	}

	finishQuizErr := p.FinishQuiz(quizID)
	if finishQuizErr != nil {
		t.Fatalf("finish a quiz before starting does not fail")
	}
}

func TestParticipant_SelectQuizAnswer_ErrosWhenNotJoinedTheQuiz(t *testing.T) {
	selectedAnswerID := inmemory.FirstAnswerID
	selectedQuestionID := inmemory.FirstQuestionID

	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	err := p.SelectQuizAnswer(quizID, selectedQuestionID, selectedAnswerID)

	if err == nil {
		t.Fatalf("expected error when providing an answer for a quiz that is not active")
	}
}

func TestParticipant_SelectQuizAnswer_StoresSelectedAnswers(t *testing.T) {
	selectedQuestionID := inmemory.FirstQuestionID
	selectedAnswerID := inmemory.FirstAnswerID

	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	err.PanicIfError(p.StartQuiz(quizID, nil))
	err.PanicIfError(p.SelectQuizAnswer(quizID, selectedQuestionID, selectedAnswerID))

	activeQuizAnswers := err.PanicIfError1(p.GetActiveQuizAnswers(quizID))

	if len(activeQuizAnswers) != 1 {
		t.Fatalf("should only contain %v answer, but contains %v", 1, len(activeQuizAnswers))
	}

	if activeQuizAnswers[0].AnswerID != selectedAnswerID {
		t.Fatalf("not expected first answer provided %v, but should be %v instead", activeQuizAnswers[0].AnswerID, selectedAnswerID)
	}
}

func TestParticipant_SelectQuizAnswer_FailsStoringAnswersForAFinishedQuiz(t *testing.T) {
	selectedQuestionID := inmemory.FirstQuestionID
	selectedAnswerID := inmemory.FirstAnswerID

	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	err.PanicIfError(p.StartQuiz(quizID, nil))
	err.PanicIfError(p.FinishQuiz(quizID))
	err := p.SelectQuizAnswer(quizID, selectedQuestionID, selectedAnswerID)

	if err == nil {
		t.Fatalf("does not fail selecting an answer for a finished quiz with id %v", quizID)
	}
}

func TestParticipant_SelectQuizAnswer_CreatesANewEventToPersist(t *testing.T) {
	selectedQuestionID := inmemory.FirstQuestionID
	selectedAnswerID := inmemory.FirstAnswerID

	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	err.PanicIfError(p.StartQuiz(quizID, nil))
	p.GetNewEventsAndUpdatePersistedVersion()

	err.PanicIfError(p.SelectQuizAnswer(quizID, selectedQuestionID, selectedAnswerID))
	events := p.GetNewEventsAndUpdatePersistedVersion()

	if len(events) != 1 {
		t.Fatalf("selecting an answer did not create a new event to persist")
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

func TestParticipant_FinishQuiz_FailsIfNotAllAnswersProvided(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	requiredQuestionIds := []string{inmemory.FirstQuestionID}

	_ = p.StartQuiz(quizID, requiredQuestionIds)
	finishQuizErr := p.FinishQuiz(quizID)

	if finishQuizErr == nil {
		t.Fatalf("could finish quiz without provided all required question answers")
	}
}

func TestParticipant_FinishQuiz(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	_ = p.StartQuiz(quizID, nil)
	finishQuizErr := p.FinishQuiz(quizID)

	if finishQuizErr != nil {
		t.Fatalf("finish quiz failed")
	}
}

func TestParticipant_FinishQuiz_failsIfFinishedTwice(t *testing.T) {
	p := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	_ = p.StartQuiz(quizID, nil)
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

	err.PanicIfError(p.StartQuiz(quizID, nil))
	err.PanicIfError(p.FinishQuiz(quizID))

	newEvents := p.GetNewEventsAndUpdatePersistedVersion()
	startedQuizEvent := newEvents[len(newEvents)-1].(event.FinishedQuiz)

	if startedQuizEvent.QuizID != quizID {
		t.Fatalf("finish event quizId does not match quiz ID '%s' != '%s'", startedQuizEvent.QuizID, quizID)
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

	_ = p.StartQuiz(quizID, nil)
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

	err.PanicIfError(p1.StartQuiz(quizID, nil))
	err.PanicIfError(p1.FinishQuiz(quizID))

	participantEvents := p1.Events

	p2 := err.PanicIfError1(participant.NewFromEvents(participantEvents, true))

	if p1.CurrentVersion != p2.CurrentVersion {
		t.Fatalf("original participant's version is different to the restored version: %d != %d", p1.CurrentVersion, p2.CurrentVersion)
	}

}

func TestParticipant_Events_createsSameEventsAfterApplyAndRestore(t *testing.T) {

	p1 := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	err.PanicIfError(p1.StartQuiz(quizID, nil))
	err.PanicIfError(p1.FinishQuiz(quizID))

	participantEvents := p1.Events

	p2 := err.PanicIfError1(participant.NewFromEvents(participantEvents, true))

	quiz2Id := newUUID()
	err.PanicIfError(p1.StartQuiz(quiz2Id, nil))
	err.PanicIfError(p2.StartQuiz(quiz2Id, nil))

	p1NewEvents := p1.GetNewEventsAndUpdatePersistedVersion()
	p1LastEvent := p1NewEvents[len(p1NewEvents)-1].(event.StartedQuiz)

	p2NewEvents := p2.GetNewEventsAndUpdatePersistedVersion()
	p2LastEvent := p2NewEvents[len(p2NewEvents)-1].(event.StartedQuiz)

	if p1LastEvent.GetVersion() != p2LastEvent.GetVersion() || p1LastEvent.QuizID != p2LastEvent.QuizID {
		t.Fatalf("original participant's new event is different to the restored version: %v != %v", p1LastEvent, p2LastEvent)
	}
}

func TestParticipant_Events_applyAndRestoresWithSameEvents(t *testing.T) {

	p1 := err.PanicIfError1(participant.New())
	quizID := err.PanicIfError1(uuid.NewRandom()).String()

	err.PanicIfError(p1.StartQuiz(quizID, nil))
	err.PanicIfError(p1.FinishQuiz(quizID))

	participantEvents := p1.Events

	p2 := err.PanicIfError1(participant.NewFromEvents(participantEvents, true))

	p1EventsAsJSON := string(err.PanicIfError1(json.Marshal(p1.Events)))
	p2EventsAsJSON := string(err.PanicIfError1(json.Marshal(p2.Events)))
	if p1EventsAsJSON != p2EventsAsJSON {
		t.Fatalf("original participant's events are different to the restored version: %s != %s", p1EventsAsJSON, p2EventsAsJSON)
	}
}

func newUUID() string {
	return err.PanicIfError1(uuid.NewRandom()).String()
}

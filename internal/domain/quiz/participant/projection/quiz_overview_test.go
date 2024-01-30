package projection_test

import (
	"learn-to-code/internal/domain/eventsource"
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/domain/quiz/participant/event"
	"learn-to-code/internal/domain/quiz/participant/projection"
	"learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/go/util/uuid"
	"testing"
)

func TestGetFinishedQuizLatestAttempt_ExistingQuiz(t *testing.T) {
	quizID := "test-quiz-id"
	events := []eventsource.Event{
		event.StartedQuiz{QuizID: quizID},
		event.FinishedQuiz{QuizID: quizID},
	}
	p := err.PanicIfError1(participant.NewFromEvents(events, true))

	qo := err.PanicIfError1(projection.NewQuizOverview(p))

	latestAttempt, err := qo.GetFinishedQuizLatestAttempt(quizID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if latestAttempt.QuizID != quizID {
		t.Fatalf("Expected quiz QuizID %s, got %s", quizID, latestAttempt.QuizID)
	}
}

func TestGetFinishedQuizLatestAttempt_NonExistingQuiz(t *testing.T) {
	p := newParticipant()

	qo := err.PanicIfError1(projection.NewQuizOverview(p))

	_, err := qo.GetFinishedQuizLatestAttempt("non-existing-quiz-id")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestGetFinishedQuizLatestAttempt_NoFinishedAttempt(t *testing.T) {
	quizID := "test-quiz-id"

	events := []eventsource.Event{
		event.StartedQuiz{QuizID: quizID},
	}
	p := err.PanicIfError1(participant.NewFromEvents(events, true))

	qo := err.PanicIfError1(projection.NewQuizOverview(p))

	_, err := qo.GetFinishedQuizLatestAttempt(quizID)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestNewQuizOverview_NoActiveQuizzes(t *testing.T) {
	quizID := "test-quiz-id"
	events := []eventsource.Event{
		event.StartedQuiz{QuizID: quizID}, event.FinishedQuiz{QuizID: quizID},
	}
	p := err.PanicIfError1(participant.NewFromEvents(events, true))

	qo := err.PanicIfError1(projection.NewQuizOverview(p))

	if len(qo.ActiveQuizzes) != 0 {
		t.Fatalf("Expected no active quizzes, found some")
	}
}

func TestNewQuizOverview_NoFinishedQuizzes(t *testing.T) {
	quizID := "test-quiz-id"
	events := []eventsource.Event{
		event.StartedQuiz{QuizID: quizID},
	}
	p := err.PanicIfError1(participant.NewFromEvents(events, true))

	qo := err.PanicIfError1(projection.NewQuizOverview(p))

	if len(qo.FinishedQuizzes) != 0 {
		t.Fatalf("Expected no finished quizzes, found some")
	}
}

func TestQuizAttemptOverview_QuizAttemptCount(t *testing.T) {
	quizID := "test-quiz-id"
	events := []eventsource.Event{
		event.StartedQuiz{QuizID: quizID}, event.FinishedQuiz{QuizID: quizID},
		event.StartedQuiz{QuizID: quizID}, event.FinishedQuiz{QuizID: quizID},
	}
	p := err.PanicIfError1(participant.NewFromEvents(events, true))

	qo := err.PanicIfError1(projection.NewQuizOverview(p))

	if len(qo.FinishedQuizzes[quizID]) != 2 {
		t.Fatalf("Expected 2 attempts for the quiz, got %d", len(qo.FinishedQuizzes[quizID]))
	}
}

func TestQuizAttemptOverview_PassIsFalseForUnfinishedQuizzes(t *testing.T) {
	quizID := "test-quiz-id"

	p := err.PanicIfError1(participant.New())
	p.StartQuiz(quizID, []string{"a"})
	p.SelectQuizAnswer(quizID, "a", "a-1", true)

	qo := err.PanicIfError1(projection.NewQuizOverview(p))

	if qo.ActiveQuizzes[quizID][0].Pass != false {
		t.Fatalf("Expected pass for not finished quizzes to be false, got %v", qo.ActiveQuizzes[quizID][0].Pass)
	}
}

func TestQuizAttemptOverview_PassIsTrueForFinishedCorrectQuizzes(t *testing.T) {
	quizID := "test-quiz-id"

	p := err.PanicIfError1(participant.New())
	err.PanicIfError(p.StartQuiz(quizID, []string{"a"}))
	err.PanicIfError(p.SelectQuizAnswer(quizID, "a", "a-1", true))
	err.PanicIfError(p.FinishQuiz(quizID))

	qo := err.PanicIfError1(projection.NewQuizOverview(p))

	if qo.FinishedQuizzes[quizID][0].Pass != true {
		t.Fatalf("Expected pass for finished and correct quizzes to be true, got %v", qo.FinishedQuizzes[quizID][0].Pass)
	}
}

func TestQuizAttemptOverview_PassIsFalseForFinishedWrongQuizzes(t *testing.T) {
	quizID := "test-quiz-id"

	p := err.PanicIfError1(participant.New())
	err.PanicIfError(p.StartQuiz(quizID, []string{"a"}))
	err.PanicIfError(p.SelectQuizAnswer(quizID, "a", "a-1", false))
	err.PanicIfError(p.FinishQuiz(quizID))

	qo := err.PanicIfError1(projection.NewQuizOverview(p))

	if qo.FinishedQuizzes[quizID][0].Pass != false {
		t.Fatalf("Expected pass for finished and wrong quizzes to be false, got %v", qo.FinishedQuizzes[quizID][0].Pass)
	}
}

func newParticipant() participant.Participant {
	return err.PanicIfError1(participant.NewParticipant(uuid.MustNewRandomAsString()))
}

package projection_test

import (
	"learn-to-code/internal/domain/quiz/participant/projection"
	"learn-to-code/internal/infrastructure/inmemory"
	"testing"
)

func TestNewQuizAttemptDetail_ErrorsForEmptyUsers(t *testing.T) {
	p := newParticipant()

	_, err := projection.NewQuizAttemptDetail(p, inmemory.QuizID, 0)

	if err == nil {
		t.Fatalf("quiz attempt detail creation returns no error for non existing quiz attempt")
	}
}

func TestNewQuizAttemptDetail_ReturnsForSingleQuestionSingleQuizUser(t *testing.T) {
	p := newParticipant()
	p.StartQuiz(inmemory.QuizID, []string{inmemory.FirstQuestionID})
	p.SelectQuizAnswer(inmemory.QuizID, inmemory.FirstQuestionID, inmemory.FirstAnswerID)

	quizAttemptDetailProjection, err := projection.NewQuizAttemptDetail(p, inmemory.QuizID, 0)

	if err != nil {
		t.Fatalf("quiz attempt detail creation errors for valid attempt")
	}

	assertQuestionAnswer(t, quizAttemptDetailProjection, inmemory.FirstQuestionID, inmemory.FirstAnswerID)
}

func TestNewQuizAttemptDetail_ReturnsForFinishedQuiz(t *testing.T) {
	p := newParticipant()

	p.StartQuiz("otherQuizId", []string{"z"})
	p.SelectQuizAnswer("otherQuizId", "z", "z-1")

	p.StartQuiz(inmemory.QuizID, []string{"a", "b", "c", "d"})
	p.SelectQuizAnswer(inmemory.QuizID, "a", "a-1")
	p.SelectQuizAnswer("otherQuizId", "otherQuestionId", inmemory.FirstAnswerID)
	p.SelectQuizAnswer(inmemory.QuizID, "b", "b-1")
	p.SelectQuizAnswer(inmemory.QuizID, "c", "c-1")
	p.SelectQuizAnswer(inmemory.QuizID, "c", "c-2")
	p.SelectQuizAnswer(inmemory.QuizID, "d", "d-3")
	p.FinishQuiz(inmemory.QuizID)

	p.StartQuiz(inmemory.QuizID, []string{"a", "b", "c", "d"})
	p.SelectQuizAnswer(inmemory.QuizID, "a", "a-2")
	p.SelectQuizAnswer("otherQuizId", "otherQuestionId", inmemory.FirstAnswerID)
	p.SelectQuizAnswer(inmemory.QuizID, "b", "b-2")
	p.SelectQuizAnswer(inmemory.QuizID, "c", "c-2")
	p.SelectQuizAnswer(inmemory.QuizID, "c", "c-3")
	p.SelectQuizAnswer(inmemory.QuizID, "d", "d-4")
	p.FinishQuiz(inmemory.QuizID)

	quizAttemptDetailProjection, err := projection.NewQuizAttemptDetail(p, inmemory.QuizID, 0)

	if err != nil {
		t.Fatalf("quiz attempt detail creation errors for valid attempt")
	}

	assertQuestionAnswer(t, quizAttemptDetailProjection, "a", "a-1")
	assertQuestionAnswer(t, quizAttemptDetailProjection, "c", "c-2")

	quizAttemptDetailProjection2, err := projection.NewQuizAttemptDetail(p, inmemory.QuizID, 1)

	if err != nil {
		t.Fatalf("quiz attempt detail creation errors for valid attempt")
	}

	assertQuestionAnswer(t, quizAttemptDetailProjection2, "a", "a-2")
	assertQuestionAnswer(t, quizAttemptDetailProjection2, "c", "c-3")
}

func assertQuestionAnswer(t *testing.T, quizAttemptDetailProjection projection.QuizAttemptDetail, questionId string, questionAnswer string) {
	providedAnswer, ok := quizAttemptDetailProjection.QuestionsWithAnswer[questionId]

	if !ok {
		t.Fatalf("missing answer %s in quiz attempt detail projection", questionId)
	}

	if providedAnswer != questionAnswer {
		t.Fatalf("wrong answer %s in quiz attempt detail projection, expected %s", providedAnswer, questionAnswer)
	}
}

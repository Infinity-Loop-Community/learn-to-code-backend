package quizattemptdetail

import (
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/infrastructure/go/util/err"
	"learn-to-code/internal/infrastructure/go/util/uuid"
	"learn-to-code/internal/infrastructure/inmemory"
	"testing"
)

func TestNewQuizAttemptDetail_ErrorsForEmptyUsers(t *testing.T) {
	p := newParticipant()

	_, err := NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, 1)

	if err == nil {
		t.Fatalf("quiz attempt detail creation returns no error for non existing quiz attempt")
	}
}

func TestNewQuizAttemptDetail_ReturnsForSingleQuestionSingleQuizUser(t *testing.T) {
	p := newParticipant()
	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{inmemory.FirstQuestionID}))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, inmemory.FirstQuestionID, inmemory.FirstAnswerID, true))

	quizAttemptDetailProjection, err := NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, 1)

	if err != nil {
		t.Fatalf("quiz attempt detail creation errors for valid attempt: %v", err)
	}

	assertQuestionAnswer(t, quizAttemptDetailProjection, inmemory.FirstQuestionID, inmemory.FirstAnswerID)
}

func TestNewQuizAttemptDetail_StartedQuiz_ReturnsOngoingAttemptState(t *testing.T) {
	p := newParticipant()
	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{inmemory.FirstQuestionID}))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, inmemory.FirstQuestionID, inmemory.FirstAnswerID, true))
	err.PanicIfError(p.FinishQuiz(inmemory.QuizIDEssentialsOfTheWeb))

	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{inmemory.FirstQuestionID}))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, inmemory.FirstQuestionID, inmemory.FirstAnswerID, true))

	quizAttemptDetailProjection := err.PanicIfError1(NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, p.GetQuizAttemptCount(inmemory.QuizIDEssentialsOfTheWeb)))

	if quizAttemptDetailProjection.AttemptStatus != AttemptStatusOngoing {
		t.Fatalf("started quiz is not ongoing")
	}
}

func TestNewQuizAttemptDetail_StartedQuiz_ReturnsValidAttemptID(t *testing.T) {
	p := newParticipant()
	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{inmemory.FirstQuestionID}))

	quizAttemptDetailProjection := err.PanicIfError1(NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, 1))

	if quizAttemptDetailProjection.AttemptID != 1 {
		t.Fatalf("started quiz has not attemptID 1, it has instead %d", quizAttemptDetailProjection.AttemptID)
	}
}

func TestNewQuizAttemptDetail_FinishedQuiz_ReturnsFinishedAttemptState(t *testing.T) {
	p := newParticipant()
	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{}))
	err.PanicIfError(p.FinishQuiz(inmemory.QuizIDEssentialsOfTheWeb))

	quizAttemptDetailProjection := err.PanicIfError1(NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, p.GetQuizAttemptCount(inmemory.QuizIDEssentialsOfTheWeb)))

	if quizAttemptDetailProjection.AttemptStatus != AttemptStatusFinished {
		t.Fatalf("finished quiz is not in finished state")
	}
}

func TestNewQuizAttemptDetail_FinishedQuiz_ReturnsQuizResultPass(t *testing.T) {
	p := newParticipant()
	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{}))
	err.PanicIfError(p.FinishQuiz(inmemory.QuizIDEssentialsOfTheWeb))

	quizAttemptDetailProjection := err.PanicIfError1(NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, p.GetQuizAttemptCount(inmemory.QuizIDEssentialsOfTheWeb)))

	if quizAttemptDetailProjection.AttemptResult.Pass != true {
		t.Fatalf("finished quiz with no questions did not pass")
	}
}

func TestNewQuizAttemptDetail_FinishedQuizWithAllWrongAnswers_ReturnsQuizResultNotPass(t *testing.T) {
	p := newParticipant()
	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{"q1"}))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q1", "a1", false))
	err.PanicIfError(p.FinishQuiz(inmemory.QuizIDEssentialsOfTheWeb))

	quizAttemptDetailProjection := err.PanicIfError1(NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, p.GetQuizAttemptCount(inmemory.QuizIDEssentialsOfTheWeb)))

	if quizAttemptDetailProjection.AttemptResult.Pass != false {
		t.Fatalf("finished quiz with only wrong answers did pass")
	}
}

func TestNewQuizAttemptDetail_FinishedQuizWithMostCorrectAnswers_ReturnsQuizResultPass(t *testing.T) {
	p := newParticipant()
	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{"q1", "q2", "q3", "q4", "q5"}))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q1", "a1", false))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q2", "a2", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q3", "a3", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q4", "a4", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q5", "a5", true))
	err.PanicIfError(p.FinishQuiz(inmemory.QuizIDEssentialsOfTheWeb))

	quizAttemptDetailProjection := err.PanicIfError1(NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, p.GetQuizAttemptCount(inmemory.QuizIDEssentialsOfTheWeb)))

	if quizAttemptDetailProjection.AttemptResult.Pass != true {
		t.Fatalf("finished quiz with mostly correct answers did not pass")
	}
}

func TestNewQuizAttemptDetail_FinishedQuizSeveralIncorrectAnswers_ReturnsQuizResultNotPass(t *testing.T) {
	p := newParticipant()
	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{"q1", "q2", "q3", "q4", "q5"}))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q1", "a1", false))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q2", "a2", false))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q3", "a3", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q4", "a4", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q5", "a5", true))
	err.PanicIfError(p.FinishQuiz(inmemory.QuizIDEssentialsOfTheWeb))

	quizAttemptDetailProjection := err.PanicIfError1(NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, p.GetQuizAttemptCount(inmemory.QuizIDEssentialsOfTheWeb)))

	if quizAttemptDetailProjection.AttemptResult.Pass != false {
		t.Fatalf("finished quiz with some incorrect answers did pass")
	}
}

func TestNewQuizAttemptDetail_FinishedQuizSeveralIncorrectAnswers_ReturnsCorrectnessRatio(t *testing.T) {
	p := newParticipant()
	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{"q1", "q2", "q3", "q4", "q5"}))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q1", "a1", false))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q2", "a2", false))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q3", "a3", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q4", "a4", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q5", "a5", true))
	err.PanicIfError(p.FinishQuiz(inmemory.QuizIDEssentialsOfTheWeb))

	quizAttemptDetailProjection := err.PanicIfError1(NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, p.GetQuizAttemptCount(inmemory.QuizIDEssentialsOfTheWeb)))

	if quizAttemptDetailProjection.AttemptResult.QuestionCorrectRatio != 0.6 {
		t.Fatalf("expected correctness ratio of 0.8 but was %f", quizAttemptDetailProjection.AttemptResult.QuestionCorrectRatio)
	}
}

func TestNewQuizAttemptDetail_FinishedQuizWithDuplicatedAnswers_ReturnsCorrectnessRatio(t *testing.T) {
	p := newParticipant()

	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{"q1", "q2", "q3"}))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q1", "a1", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q2", "a2", false))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q2", "a2", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q3", "a3", true))
	err.PanicIfError(p.FinishQuiz(inmemory.QuizIDEssentialsOfTheWeb))

	quizAttemptDetailProjection := err.PanicIfError1(NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, p.GetQuizAttemptCount(inmemory.QuizIDEssentialsOfTheWeb)))

	if quizAttemptDetailProjection.AttemptResult.QuestionCorrectRatio != 1 {
		t.Fatalf("expected correctness ratio of 1 but was %f", quizAttemptDetailProjection.AttemptResult.QuestionCorrectRatio)
	}
}

func TestNewQuizAttemptDetail_FinishedQuizSeveralIncorrectAnswers_ReturnsComparedToAverageRatio(t *testing.T) {
	p := newParticipant()
	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{"q1", "q2", "q3", "q4", "q5"}))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q1", "a1", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q2", "a2", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q3", "a3", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q4", "a4", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q5", "a5", true))
	err.PanicIfError(p.FinishQuiz(inmemory.QuizIDEssentialsOfTheWeb))

	quizAttemptDetailProjection := err.PanicIfError1(NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, p.GetQuizAttemptCount(inmemory.QuizIDEssentialsOfTheWeb)))

	// the minimum duration for the calculation is 1m, and we have 10 questions for each question 1m
	// means 1/10 - 1 = -0.9 * 100 = -90
	if quizAttemptDetailProjection.AttemptResult.ComparedToTimeAveragePercentage != -90 {
		t.Fatalf("expected compared to average time percentage of -90 percentage but was %d", quizAttemptDetailProjection.AttemptResult.ComparedToTimeAveragePercentage)
	}
}

func TestNewQuizAttemptDetail_FinishedQuiz_ReturnsTimeTaken(t *testing.T) {
	p := newParticipant()
	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{"q1", "q2"}))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q1", "a1", false))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q2", "a2", false))
	err.PanicIfError(p.FinishQuiz(inmemory.QuizIDEssentialsOfTheWeb))

	quizAttemptDetailProjection := err.PanicIfError1(NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, p.GetQuizAttemptCount(inmemory.QuizIDEssentialsOfTheWeb)))

	if quizAttemptDetailProjection.AttemptResult.TimeTakenMins != 1 {
		t.Fatalf("expected 1 time taken for the quiz")
	}
}

func TestNewQuizAttemptDetail_SingleFinishedQuiz_ReturnsCompareLastTryPercentage(t *testing.T) {
	p := newParticipant()
	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{"q1", "q2"}))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q1", "a1", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q2", "a2", true))
	err.PanicIfError(p.FinishQuiz(inmemory.QuizIDEssentialsOfTheWeb))

	quizAttemptDetailProjection := err.PanicIfError1(NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, p.GetQuizAttemptCount(inmemory.QuizIDEssentialsOfTheWeb)))

	if quizAttemptDetailProjection.AttemptResult.ComparedToCorrectRatioLastTryPercentage != 100 {
		t.Fatalf("expected 100 percent as compared to last try percentage because there was no last try, but was %d", quizAttemptDetailProjection.AttemptResult.ComparedToCorrectRatioLastTryPercentage)
	}
}

func TestNewQuizAttemptDetail_TwoFinishedQuizzes_ReturnsCompareLastTryPercentage(t *testing.T) {
	p := newParticipant()

	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{"q1", "q2"}))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q1", "a1", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q2", "a2", false))
	err.PanicIfError(p.FinishQuiz(inmemory.QuizIDEssentialsOfTheWeb))

	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{"q1", "q2"}))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q1", "a1", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "q2", "a3", true))
	err.PanicIfError(p.FinishQuiz(inmemory.QuizIDEssentialsOfTheWeb))

	quizAttemptDetailProjection := err.PanicIfError1(NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, p.GetQuizAttemptCount(inmemory.QuizIDEssentialsOfTheWeb)))

	if quizAttemptDetailProjection.AttemptResult.ComparedToCorrectRatioLastTryPercentage != 50 {
		t.Fatalf("expected 50 percent as compared to last try, but was %d", quizAttemptDetailProjection.AttemptResult.ComparedToCorrectRatioLastTryPercentage)
	}
}

func TestNewQuizAttemptDetail_ReturnsForFinishedQuiz(t *testing.T) {
	p := newParticipant()

	err.PanicIfError(p.StartQuiz("otherQuizId", []string{"z"}))
	err.PanicIfError(p.SelectQuizAnswer("otherQuizId", "z", "z-1", true))

	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{"a", "b", "c", "d"}))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "a", "a-1", true))
	err.PanicIfError(p.SelectQuizAnswer("otherQuizId", "otherQuestionId", inmemory.FirstAnswerID, true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "b", "b-1", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "c", "c-1", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "c", "c-2", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "d", "d-3", true))
	err.PanicIfError(p.FinishQuiz(inmemory.QuizIDEssentialsOfTheWeb))

	err.PanicIfError(p.StartQuiz(inmemory.QuizIDEssentialsOfTheWeb, []string{"a", "b", "c", "d"}))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "a", "a-2", true))
	err.PanicIfError(p.SelectQuizAnswer("otherQuizId", "otherQuestionId", inmemory.FirstAnswerID, true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "b", "b-2", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "c", "c-2", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "c", "c-3", true))
	err.PanicIfError(p.SelectQuizAnswer(inmemory.QuizIDEssentialsOfTheWeb, "d", "d-4", true))
	err.PanicIfError(p.FinishQuiz(inmemory.QuizIDEssentialsOfTheWeb))

	quizAttemptDetailProjection, err := NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, 1)

	if err != nil {
		t.Fatalf("quiz attempt detail creation errors for valid attempt")
	}
	if quizAttemptDetailProjection.AttemptID != 1 {
		t.Fatalf("invalid attampt QuizID in projection, requested id 1, but received id %d", quizAttemptDetailProjection.AttemptID)
	}

	assertQuestionAnswer(t, quizAttemptDetailProjection, "a", "a-1")
	assertQuestionAnswer(t, quizAttemptDetailProjection, "c", "c-2")

	quizAttemptDetailProjection2, err := NewQuizAttemptDetail(p, inmemory.QuizIDEssentialsOfTheWeb, 2)

	if err != nil {
		t.Fatalf("quiz attempt detail creation errors for valid attempt")
	}
	if quizAttemptDetailProjection2.AttemptID != 2 {
		t.Fatalf("invalid attampt QuizID in projection, requested id 2, but received id %d", quizAttemptDetailProjection2.AttemptID)
	}

	assertQuestionAnswer(t, quizAttemptDetailProjection2, "a", "a-2")
	assertQuestionAnswer(t, quizAttemptDetailProjection2, "c", "c-3")
}

func assertQuestionAnswer(t *testing.T, quizAttemptDetailProjection QuizAttemptDetail, questionID string, questionAnswer string) {
	providedAnswer, ok := quizAttemptDetailProjection.QuestionsWithAnswer[questionID]

	if !ok {
		t.Fatalf("missing answer %s in quiz attempt detail projection", questionID)
	}

	if providedAnswer != questionAnswer {
		t.Fatalf("wrong answer %s in quiz attempt detail projection, expected %s", providedAnswer, questionAnswer)
	}
}

func newParticipant() participant.Participant {
	return err.PanicIfError1(participant.NewParticipant(uuid.MustNewRandomAsString()))
}

package command

func NewFinishQuizData(quizID string) FinishQuiz {
	return FinishQuiz{
		QuizID: quizID,
	}
}

type FinishQuiz struct {
	QuizID string
}

const FinishQuizCommandType = "FinishQuiz"

package data

func NewStartQuizData(quizID string) StartQuiz {
	return StartQuiz{
		QuizID: quizID,
	}
}

type StartQuiz struct {
	QuizID string
}

const StartQuizCommandType = "StartQuiz"

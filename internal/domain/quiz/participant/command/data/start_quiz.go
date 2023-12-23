package data

func NewStartQuizData(quizID string, requiredQuestionsAnswered []string) StartQuiz {
	return StartQuiz{
		QuizID:                    quizID,
		RequiredQuestionsAnswered: requiredQuestionsAnswered,
	}
}

type StartQuiz struct {
	QuizID                    string
	RequiredQuestionsAnswered []string
}

const StartQuizCommandType = "StartQuiz"

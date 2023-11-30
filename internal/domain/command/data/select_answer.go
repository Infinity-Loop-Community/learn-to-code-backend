package data

func NewSelectAnswerData(quizID string, questionID string, answerID string) SelectAnswer {
	return SelectAnswer{
		QuizID:     quizID,
		QuestionID: questionID,
		AnswerID:   answerID,
	}
}

type SelectAnswer struct {
	QuizID     string
	QuestionID string
	AnswerID   string
}

const SelectAnswerCommandType = "SelectAnswer"

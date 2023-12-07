package responseobject

type QuizAttemptDetail struct {
	QuestionsWithAnswer map[string]string `json:"questionsWithAnswer"`
}

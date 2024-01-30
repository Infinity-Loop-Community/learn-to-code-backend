package responseobject

type QuizAttemptOverview struct {
	AttemptID           int               `json:"attemptId"`
	QuizID              string            `json:"quizId"`
	QuestionsWithAnswer map[string]string `json:"questionsWithAnswer"`
	Pass                bool              `json:"pass"`
}

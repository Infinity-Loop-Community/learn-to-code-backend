package projection

type QuizAttemptOverview struct {
	QuizID              string
	AttemptID           int
	QuestionsWithAnswer map[string]string
}

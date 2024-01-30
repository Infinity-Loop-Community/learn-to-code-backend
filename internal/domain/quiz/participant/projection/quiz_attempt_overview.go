package projection

type QuizAttemptOverview struct {
	QuizID              string
	AttemptID           int
	Pass                bool
	QuestionsWithAnswer map[string]string
}

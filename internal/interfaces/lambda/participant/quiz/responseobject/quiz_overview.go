package responseobject

type QuizOverview struct {
	ActiveQuizzes   map[string][]QuizAttemptOverview `json:"activeQuizzes"`
	FinishedQuizzes map[string][]QuizAttemptOverview `json:"finishedQuizzes"`
}

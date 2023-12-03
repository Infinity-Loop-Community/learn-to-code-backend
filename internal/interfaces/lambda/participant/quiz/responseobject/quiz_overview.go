package responseobject

type QuizOverview struct {
	ActiveQuizzes   []string `json:"activeQuizzes"`
	FinishedQuizzes []string `json:"finishedQuizzes"`
}

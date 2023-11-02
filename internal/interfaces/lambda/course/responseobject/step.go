package responseobject

type Step struct {
	ID      string     `json:"id"`
	Quizzes []StepQuiz `json:"quizzes"`
	Name    string     `json:"name"`
}

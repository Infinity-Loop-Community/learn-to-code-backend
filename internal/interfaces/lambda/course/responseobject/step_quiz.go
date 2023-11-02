package responseobject

type StepQuiz struct {
	ID        string         `json:"id"`
	Questions []QuizQuestion `json:"questions"`
}

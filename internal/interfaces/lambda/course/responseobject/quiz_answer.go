package responseobject

type QuizAnswer struct {
	ID          string `json:"id"`
	Text        string `json:"text"`
	IsCorrect   bool   `json:"isCorrect"`
	Description string `json:"description"`
}

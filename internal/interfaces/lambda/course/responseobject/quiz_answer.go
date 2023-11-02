package responseobject

type QuizAnswer struct {
	Text        string `json:"text"`
	IsCorrect   bool   `json:"isCorrect"`
	Description string `json:"description"`
}

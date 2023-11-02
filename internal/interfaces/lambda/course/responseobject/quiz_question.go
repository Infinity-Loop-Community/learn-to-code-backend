package responseobject

type QuizQuestion struct {
	Text       string       `json:"text"`
	Difficulty string       `json:"difficulty"`
	Answers    []QuizAnswer `json:"answers"`
}

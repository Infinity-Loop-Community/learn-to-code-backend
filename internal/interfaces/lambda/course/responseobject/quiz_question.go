package responseobject

type QuizQuestion struct {
	Text        string       `json:"text"`
	Difficulty  string       `json:"difficulty"`
	Answers     []QuizAnswer `json:"answers"`
	Rating      float32      `json:"rating"`
	RatingCount int          `json:"ratingCount"`
}

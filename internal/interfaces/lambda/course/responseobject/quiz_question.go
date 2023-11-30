package responseobject

type QuizQuestion struct {
	ID          string       `json:"id"`
	Text        string       `json:"text"`
	Difficulty  string       `json:"difficulty"`
	Answers     []QuizAnswer `json:"answers"`
	Rating      float32      `json:"rating"`
	RatingCount int          `json:"ratingCount"`
}

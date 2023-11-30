package course

type QuizQuestion struct {
	ID          string
	Text        string
	Difficulty  string
	Answers     []QuizAnswer
	Rating      float64
	RatingCount int
}

package course

type QuizQuestion struct {
	Text        string
	Difficulty  string
	Answers     []QuizAnswer
	Rating      float64
	RatingCount int
}

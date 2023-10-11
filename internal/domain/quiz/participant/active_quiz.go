package participant

type activeQuiz struct {
	ID              string
	providedAnswers []providedAnswer
	completed       bool
}

func (q activeQuiz) IsOngoing() bool {
	return !q.completed
}

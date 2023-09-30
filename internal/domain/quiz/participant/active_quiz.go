package participant

type activeQuiz struct {
	Id              string
	providedAnswers []providedAnswer
	completed       bool
}

func (q activeQuiz) IsOngoing() bool {
	return q.completed == false
}

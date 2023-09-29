package participant

type activeQuiz struct {
	id              string
	providedAnswers []providedAnswer
	completed       bool
}

func (q activeQuiz) isOngoing() bool {
	return q.completed == false
}

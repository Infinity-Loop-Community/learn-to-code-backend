package participant

type quizAttempt struct {
	ID                        string
	providedAnswers           []ProvidedAnswer
	completed                 bool
	requiredQuestionsAnswered []string
}

func (q quizAttempt) IsOngoing() bool {
	return !q.completed
}

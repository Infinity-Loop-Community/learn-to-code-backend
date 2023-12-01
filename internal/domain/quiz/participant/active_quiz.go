package participant

type activeQuiz struct {
	ID                        string
	providedAnswers           []ProvidedAnswer
	completed                 bool
	requiredQuestionsAnswered []string
}

func (q activeQuiz) IsOngoing() bool {
	return !q.completed
}

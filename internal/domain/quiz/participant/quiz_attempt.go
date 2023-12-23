package participant

import "fmt"

type quizAttempt struct {
	ID                        string
	providedAnswers           []ProvidedAnswer
	completed                 bool
	requiredQuestionsAnswered []string
}

func (q quizAttempt) IsOngoing() bool {
	return !q.completed
}

func (q quizAttempt) checkFinishAttemptValidity() error {
	err := q.checkAllQuestionsProvidedValidity()
	if err != nil {
		return err
	}

	if q.completed {
		return fmt.Errorf("Quiz %v already finished", q.ID)
	}

	return nil
}

func (q quizAttempt) checkAllQuestionsProvidedValidity() error {
	providedQuestionsLookupTable := map[string]bool{}

	for _, answer := range q.providedAnswers {
		providedQuestionsLookupTable[answer.QuestionID] = true
	}

	allAnswersProvided := true
	missingQuestionIds := []string{}
	for _, requiredQuestionAnswered := range q.requiredQuestionsAnswered {
		_, ok := providedQuestionsLookupTable[requiredQuestionAnswered]
		if !ok {
			allAnswersProvided = false
			missingQuestionIds = append(missingQuestionIds, requiredQuestionAnswered)
		}
	}

	if !allAnswersProvided {
		return fmt.Errorf("not all answers provided, the answer for the following question ids are missing: %v", missingQuestionIds)
	}
	return nil
}

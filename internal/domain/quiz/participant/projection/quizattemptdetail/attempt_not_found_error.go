package quizattemptdetail

import "fmt"

type AttemptNotFoundError struct {
	AttemptID    int
	AttemptCount int
	QuizID       string
}

func (a AttemptNotFoundError) Error() string {
	return fmt.Sprintf("quiz attempt id %v does not exist, there are only %v attempts for quizId %v", a.AttemptID, a.AttemptCount, a.QuizID)
}

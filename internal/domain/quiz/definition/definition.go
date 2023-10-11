package definition

import "fmt"

type Definition struct {
	ID        string
	Questions []Question
}

func (d Definition) IsComplete() bool {

	var errors []error

	for _, q := range d.Questions {
		questionIsComplete := len(q.PossibleAnswers) > 1

		if !questionIsComplete {
			errors = append(errors, fmt.Errorf("not sufficient answers provided, requires at least 2 answers"))
		}
	}

	return len(errors) == 0
}

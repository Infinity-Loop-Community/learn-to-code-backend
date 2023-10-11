package definition

import (
	"testing"

	"github.com/google/uuid"
)

func TestQuizAnswersMissing(t *testing.T) {
	quiz := Definition{
		ID: "QuizDefinition Text",
		Questions: []Question{
			{
				ID:   newID(t),
				Text: "Question1 Text",
				PossibleAnswers: []Answer{
					{
						ID:          newID(t),
						IsCorrect:   false,
						Explanation: "Q1: This is not the correct answer",
					},
				},
			},
		},
	}

	if quiz.IsComplete() == true {
		t.Fatalf("only %d question, expected to be not complete", len(quiz.Questions))
	}
}

func TestQuizIsComplete(t *testing.T) {
	quiz := Definition{
		ID: "QuizDefinition Text",
		Questions: []Question{
			{
				ID:   newID(t),
				Text: "Question1 Text",
				PossibleAnswers: []Answer{
					{
						ID:          newID(t),
						IsCorrect:   false,
						Explanation: "Q1: This is not the correct answer",
					},
					{
						ID:          newID(t),
						IsCorrect:   false,
						Explanation: "Q1: This is the correct answer",
					},
				},
			},
		},
	}

	if quiz.IsComplete() == false {
		t.Fatalf("%d questions, expected to be complete", len(quiz.Questions))
	}
}

func newID(t *testing.T) string {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		t.Fatalf("UUID generation failed")
	}

	return newUUID.String()
}

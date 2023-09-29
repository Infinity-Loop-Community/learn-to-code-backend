package definition

import (
	"github.com/google/uuid"
	"testing"
)

func TestQuizAnswersMissing(t *testing.T) {
	quiz := Definition{
		Id: "QuizDefinition Text",
		Questions: []Question{
			{
				Id:   newId(t),
				Text: "Question1 Text",
				PossibleAnswers: []Answer{
					Answer{
						Id:          newId(t),
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
		Id: "QuizDefinition Text",
		Questions: []Question{
			{
				Id:   newId(t),
				Text: "Question1 Text",
				PossibleAnswers: []Answer{
					{
						Id:          newId(t),
						IsCorrect:   false,
						Explanation: "Q1: This is not the correct answer",
					},
					{
						Id:          newId(t),
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

func newId(t *testing.T) string {
	newUuid, err := uuid.NewRandom()
	if err != nil {
		t.Fatalf("UUID generation failed")
	}

	return newUuid.String()
}

package command

import (
	"fmt"
	"learn-to-code/internal/domain/command/data"
	"learn-to-code/internal/domain/quiz/participant"

	"github.com/mitchellh/mapstructure"
)

type ParticipantCommandApplier struct {
}

func NewParticipantCommandApplier() *ParticipantCommandApplier {
	return &ParticipantCommandApplier{}
}

func (m *ParticipantCommandApplier) ApplyCommand(command Command, p participant.Participant) (participant.Participant, error) {
	switch command.Type {
	case data.StartQuizCommandType:
		startQuiz := data.StartQuiz{}
		mapstructure.Decode(command.Data, &startQuiz)

		p.StartQuiz(startQuiz.QuizID)

	case data.SelectAnswerCommandType:
		selectAnswerData := data.SelectAnswer{}
		mapstructure.Decode(command.Data, &selectAnswerData)
		p.SelectQuizAnswer(selectAnswerData.QuizID, selectAnswerData.QuestionID, selectAnswerData.AnswerID)

	default:
		return p, fmt.Errorf("unknown command type '%s'", command.Type)
	}

	return p, nil
}

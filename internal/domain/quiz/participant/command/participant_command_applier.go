package command

import (
	"fmt"
	"learn-to-code/internal/domain/quiz/participant"

	"github.com/mitchellh/mapstructure"
)

type ParticipantCommandApplier struct {
}

func NewParticipantCommandApplier() *ParticipantCommandApplier {
	return &ParticipantCommandApplier{}
}

func (m *ParticipantCommandApplier) ApplyCommand(command Command, participant participant.Participant) (participant.Participant, error) {
	var err error = nil

	switch command.Type {
	case StartQuizCommandType:
		startQuizCommand := StartQuiz{}
		mapstructure.Decode(command.Data, &startQuizCommand)

		err = participant.StartQuiz(startQuizCommand.QuizID, startQuizCommand.RequiredQuestionsAnswered)

	case SelectAnswerCommandType:
		selectAnswerCommand := SelectAnswer{}
		mapstructure.Decode(command.Data, &selectAnswerCommand)
		err = participant.SelectQuizAnswer(selectAnswerCommand.QuizID, selectAnswerCommand.QuestionID, selectAnswerCommand.AnswerID)

	case FinishQuizCommandType:
		finishQuizCommand := FinishQuiz{}
		mapstructure.Decode(command.Data, &finishQuizCommand)
		err = participant.FinishQuiz(finishQuizCommand.QuizID)

	default:
		return participant, fmt.Errorf("unknown command type '%s'", command.Type)
	}

	return participant, err
}

package command

import (
	"fmt"
	"learn-to-code/internal/domain/quiz/participant"
	"learn-to-code/internal/domain/quiz/participant/command/data"

	"github.com/mitchellh/mapstructure"
)

type ParticipantCommandApplier struct {
}

func NewParticipantCommandApplier() *ParticipantCommandApplier {
	return &ParticipantCommandApplier{}
}

func (m *ParticipantCommandApplier) ApplyCommand(command Command, p participant.Participant) (participant.Participant, error) {
	var err error = nil

	switch command.Type {
	case data.StartQuizCommandType:
		startQuiz := data.StartQuiz{}
		mapstructure.Decode(command.Data, &startQuiz)

		err = p.StartQuiz(startQuiz.QuizID, startQuiz.RequiredQuestionsAnswered)

	case data.SelectAnswerCommandType:
		selectAnswerData := data.SelectAnswer{}
		mapstructure.Decode(command.Data, &selectAnswerData)
		err = p.SelectQuizAnswer(selectAnswerData.QuizID, selectAnswerData.QuestionID, selectAnswerData.AnswerID)

	case data.FinishQuizCommandType:
		finishQuizData := data.FinishQuiz{}
		mapstructure.Decode(command.Data, &finishQuizData)
		err = p.FinishQuiz(finishQuizData.QuizID)

	default:
		return p, fmt.Errorf("unknown command type '%s'", command.Type)
	}

	return p, err
}

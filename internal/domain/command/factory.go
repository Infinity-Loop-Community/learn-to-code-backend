package command

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

type Factory struct {
}

func NewCommandFactory() *Factory {
	return &Factory{}
}

func (f *Factory) CreateStartQuizCommand(quizID string, requiredQuestionsAnswered []string) Command {
	return NewCommand(StartQuizCommandType, NewStartQuizData(quizID, requiredQuestionsAnswered), time.Now())
}

func (f *Factory) CreateSelectAnswerCommand(quizID string, questionID string, answerID string) Command {
	return NewCommand(SelectAnswerCommandType, NewSelectAnswerData(quizID, questionID, answerID), time.Now())
}

func (f *Factory) CreateFinishQuizCommand(quizID string) Command {
	return NewCommand(FinishQuizCommandType, NewFinishQuizData(quizID), time.Now())
}

func DecodeCommand[T any](data any, result T) (T, error) {
	err := decode(data, result)
	return result, err
}

func decode(data any, result any) error {
	config := &mapstructure.DecoderConfig{
		ErrorUnset: true,
		Result:     result,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	err = decoder.Decode(data)

	return err
}

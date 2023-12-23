package command

import (
	"learn-to-code/internal/domain/command/data"
	"time"
)

type Factory struct {
}

func NewCommandFactory() *Factory {
	return &Factory{}
}

func (f *Factory) CreateStartQuizCommand(quizID string, requiredQuestionsAnswered []string) Command {
	return NewCommand(data.StartQuizCommandType, data.NewStartQuizData(quizID, requiredQuestionsAnswered), time.Now())
}

func (f *Factory) CreateSelectAnswerCommand(quizID string, questionID string, answerID string) Command {
	return NewCommand(data.SelectAnswerCommandType, data.NewSelectAnswerData(quizID, questionID, answerID), time.Now())
}

func (f *Factory) CreateFinishQuizCommand(quizID string) Command {
	return NewCommand(data.FinishQuizCommandType, data.NewFinishQuizData(quizID), time.Now())
}

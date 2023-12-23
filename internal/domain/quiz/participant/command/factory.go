package command

import (
	"time"
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

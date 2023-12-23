package event

import (
	"learn-to-code/internal/domain/eventsource"
	"reflect"
)

type SelectedAnswer struct {
	QuizID     string
	QuestionID string
	AnswerID   string
	IsCorrect  bool
	eventsource.EventBase
}

var SelectedAnswerTypeName = reflect.TypeOf(SelectedAnswer{}).Name()

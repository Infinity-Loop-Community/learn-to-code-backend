package event

import (
	"learn-to-code/internal/domain/eventsource"
	"reflect"
)

type FinishedQuiz struct {
	QuizID string
	eventsource.EventBase
	Pass bool
}

var FinishedQuizTypeName = reflect.TypeOf(FinishedQuiz{}).Name()

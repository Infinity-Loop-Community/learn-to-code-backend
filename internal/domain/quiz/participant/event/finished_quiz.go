package event

import (
	"learn-to-code/internal/domain/eventsource"
	"reflect"
)

type FinishedQuiz struct {
	eventsource.EventBase
}

var FinishedQuizTypeName = reflect.TypeOf(FinishedQuiz{}).Name()

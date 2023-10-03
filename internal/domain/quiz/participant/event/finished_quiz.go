package event

import (
	"hello-world/internal/domain/eventsource"
	"reflect"
)

type FinishedQuiz struct {
	eventsource.EventBase
}

var FinishedQuizTypeName = reflect.TypeOf(FinishedQuiz{}).Name()

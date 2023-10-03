package event

import (
	"hello-world/internal/domain/eventsource"
	"reflect"
)

type StartedQuiz struct {
	eventsource.EventBase
}

var StartedQuizTypeName = reflect.TypeOf(StartedQuiz{}).Name()

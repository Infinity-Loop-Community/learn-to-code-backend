package event

import (
	"learn-to-code/internal/domain/eventsource"
	"reflect"
)

type StartedQuiz struct {
	eventsource.EventBase
}

var StartedQuizTypeName = reflect.TypeOf(StartedQuiz{}).Name()

package event

import (
	"hello-world/internal/domain/eventsource"
	"reflect"
)

type JoinedQuiz struct {
	eventsource.EventBase
}

var JoinedQuizTypeName = reflect.TypeOf(JoinedQuiz{}).Name()

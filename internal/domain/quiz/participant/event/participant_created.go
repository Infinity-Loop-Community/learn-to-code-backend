package event

import (
	"hello-world/internal/domain/eventsource"
	"reflect"
)

type ParticipantCreated struct {
	eventsource.EventBase
}

var ParticipantCreatedTypeName = reflect.TypeOf(ParticipantCreated{}).Name()

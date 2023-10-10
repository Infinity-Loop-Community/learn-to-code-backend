package event

import (
	"learn-to-code/internal/domain/eventsource"
	"reflect"
)

type ParticipantCreated struct {
	eventsource.EventBase
}

var ParticipantCreatedTypeName = reflect.TypeOf(ParticipantCreated{}).Name()

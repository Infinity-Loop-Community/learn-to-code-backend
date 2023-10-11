package eventsource

import "time"

type Event interface {
	GetID() string
	GetVersion() uint
	GetCreatedAt() time.Time
}

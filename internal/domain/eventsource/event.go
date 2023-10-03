package eventsource

import "time"

type Event interface {
	GetId() string
	GetVersion() uint
	GetCreatedAt() time.Time
}

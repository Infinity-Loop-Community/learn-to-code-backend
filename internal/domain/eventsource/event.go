package eventsource

import "time"

type Event interface {
	GetID() string
	GetVersion() uint
	GetAggregateID() string
	GetCreatedAt() time.Time
}

package eventsource

import "time"

type Event interface {
	GetVersion() uint
	GetAggregateID() string
	GetCreatedAt() time.Time
}

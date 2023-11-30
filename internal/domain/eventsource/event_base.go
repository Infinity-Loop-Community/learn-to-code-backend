package eventsource

import "time"

type EventBase struct {
	AggregateID string
	Version     uint
	CreatedAt   time.Time
}

func (a EventBase) GetAggregateID() string {
	return a.AggregateID
}

func (a EventBase) GetVersion() uint {
	return a.Version
}

func (a EventBase) GetCreatedAt() time.Time {
	return a.CreatedAt
}

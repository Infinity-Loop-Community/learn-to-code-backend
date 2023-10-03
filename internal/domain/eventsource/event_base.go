package eventsource

import "time"

type EventBase struct {
	Id        string
	Version   uint
	CreatedAt time.Time
}

func (a EventBase) GetId() string {
	return a.Id
}

func (a EventBase) GetVersion() uint {
	return a.Version
}

func (a EventBase) GetCreatedAt() time.Time {
	return a.CreatedAt
}

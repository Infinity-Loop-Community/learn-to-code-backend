package eventsource

import "time"

type Event interface {
	// GetVersion returns the version of the event, which is essential for maintaining the sequence
	// of events and the entity's state over time.
	GetVersion() uint

	// GetAggregateID provides the identifier of the aggregate to which the event belongs,
	// linking the event to its respective entity.
	GetAggregateID() string

	// GetCreatedAt the timestamp of when the event was created, providing a temporal context to
	// the changes represented by the event.
	GetCreatedAt() time.Time
}

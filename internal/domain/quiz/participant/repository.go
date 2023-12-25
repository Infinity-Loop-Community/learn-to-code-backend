package participant

import (
	"learn-to-code/internal/domain/eventsource"
)

// The Repository needs transactional safety to ensure proper functioning of event sourcing.
// This ensures that events are inserted in the correct order. In traditional relational databases,
// you can achieve this by using transactions. In NoSQL databases like DynamoDB,
// you can utilize TransactionalWrites along with condition checks, such as an event counter,
// to maintain transactional integrity.
type Repository interface {
	StoreEvents(participantID string, events []eventsource.Event) error

	// FindByID retrieves a participant by ID from the repository or creates an empty one if not exists.
	FindOrCreateByID(participantID string) (Participant, error)

	FindEventsByParticipantID(pariticipantID string) ([]eventsource.Event, error)
}

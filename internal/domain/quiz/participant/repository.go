package participant

import (
	"errors"
	"learn-to-code/internal/domain/eventsource"
)

var ErrNotFound = errors.New("element not found")

// The Repository needs transactional safety to ensure proper functioning of event sourcing.
// This ensures that events are inserted in the correct order. In traditional relational databases,
// you can achieve this by using transactions. In NoSQL databases like DynamoDB,
// you can utilize TransactionalWrites along with condition checks, such as an event counter,
// to maintain transactional integrity.
type Repository interface {
	AppendEvents(participantId string, events []eventsource.Event) error

	// FindById retrieves a participant by ID from the repository.
	// It returns the Participant object and an error.
	// If the user is not found, a static error ErrNotFound must be returned.
	FindById(id string) (Participant, error)
}

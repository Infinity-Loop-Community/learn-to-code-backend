package inmemory

import (
	. "hello-world/internal/domain/quiz/participant"
	"hello-world/internal/domain/quiz/participant/event"
)

type ParticipantRepository struct {
	data map[string][]event.Event
}

func NewParticipantRepository() *ParticipantRepository {
	return &ParticipantRepository{
		data: make(map[string][]event.Event),
	}
}

func (r *ParticipantRepository) AppendEvent(id string, e event.Event) error {
	r.data[id] = append(r.data[id], e)
	return nil
}

func (r *ParticipantRepository) FindById(id string) (Participant, error) {
	events, ok := r.data[id]

	if ok != true {
		return Participant{}, ErrNotFound
	}

	p, err := NewFromEvents(id, events)

	return p, err
}

package eventsource

// AggregateRoot acts as the main entity around which events are centered, encapsulating key
// business logic and state changes.
type AggregateRoot struct {

	// persistedVersion keeps track of the last persisted version of the entity, ensuring data
	// integrity and consistency during the event storing process.
	persistedVersion uint

	// currentVersion represents the current version of the entity, reflecting all the changes made
	// by the events that have occurred since the last persistence.
	currentVersion uint

	// events is a collection of events that have been applied to the entity but not yet persisted.
	// These events encapsulate the changes that have occurred in the entity's state.
	events []Event
}

func (a *AggregateRoot) GetPersistedVerstion() uint {
	return a.persistedVersion
}

func (a *AggregateRoot) IncrementPersistedVerstion() {
	a.persistedVersion++
}

func (a *AggregateRoot) GetCurrentVersion() uint {
	return a.currentVersion
}

func (a *AggregateRoot) IncremenCurrentVerstion() {
	a.currentVersion++
}

func (a *AggregateRoot) GetEvents() []Event {
	return a.events
}

func (a *AggregateRoot) AppendEvent(event Event, isPersisted bool) {

	a.IncremenCurrentVerstion()

	if isPersisted && (a.GetCurrentVersion()-1) == a.GetPersistedVerstion() {
		a.IncrementPersistedVerstion()
	}

	a.events = append(a.events, event)
}

func (a *AggregateRoot) GetNewEventsAndUpdatePersistedVersion() []Event {
	var newEvents []Event

	for _, event := range a.events {
		if event.GetVersion() >= a.persistedVersion {
			newEvents = append(newEvents, event)
		}
	}

	if int(a.persistedVersion)+len(newEvents) == int(a.currentVersion) {
		a.persistedVersion = a.currentVersion
	} else {
		panic("aggregate version count and tracking is out of sync")
	}

	return newEvents
}

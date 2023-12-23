package eventsource

type AggregateRoot struct {
	persistedVersion uint
	currentVersion   uint
	events           []Event
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

func (a *AggregateRoot) AppendEvent(event Event) {
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

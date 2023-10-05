package eventsource

type AggregateRoot struct {
	PersistedVersion uint
	CurrentVersion   uint
	Events           []Event
}

func (a *AggregateRoot) GetNewEventsAndUpdatePersistedVersion() []Event {
	var newEvents []Event

	for _, event := range a.Events {
		if event.GetVersion() >= a.PersistedVersion {
			newEvents = append(newEvents, event)
		}
	}

	if int(a.PersistedVersion)+len(newEvents) == int(a.CurrentVersion) {
		a.PersistedVersion = a.CurrentVersion
	} else {
		panic("aggregate version count and tracking is out of sync")
	}

	return newEvents
}

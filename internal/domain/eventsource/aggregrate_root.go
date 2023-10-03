package eventsource

type AggregateRoot struct {
	PersistedVersion uint
	CurrentVersion   uint
	Events           []Event
}

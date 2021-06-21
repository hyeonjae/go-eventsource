package go_eventsource

type Aggregate interface {
	GetID() string
	ListEvents() []Event
	Apply(event Event)
}

package go_eventsource

type Event interface {
	Type() string
}

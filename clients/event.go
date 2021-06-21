package clients

import "time"

type Event struct {
	AggregateId string    `bson:"aggregate_id,omitempty"`
	Timestamp   time.Time `bson:"timestamp,omitempty"`
	EventType   string    `bson:"event_name,omitempty"`
	Name        string    `bson:"name,omitempty"`
	Email       string    `bson:"email,omitempty"`
}

func NewEvent(aggregateId string, timestamp time.Time, name, email string) *Event {
	return &Event{
		AggregateId: aggregateId,
		Timestamp:   timestamp,
		EventType:   "EnrollClient",
		Name:        name,
		Email:       email,
	}
}

func (e Event) Type() string {
	return e.EventType
}

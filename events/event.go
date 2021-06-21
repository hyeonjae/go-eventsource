package events

import "time"

type Event struct {
	AggregateId string    `bson:"aggregate_id,omitempty"`
	Timestamp   time.Time `bson:"timestamp,omitempty"`
	EventType   string    `bson:"event_name,omitempty"`
	Name        string    `bson:"name,omitempty"`
	Email       string    `bson:"email,omitempty"`
}

func (e Event) Type() string {
	return e.EventType
}

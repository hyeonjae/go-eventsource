package events

import "time"

type UpdateClient struct {
	Event
}

func NewUpdateClient(aggregateId string, timestamp time.Time, name, email string) *UpdateClient {
	return &UpdateClient{
		Event: Event{
			AggregateId: aggregateId,
			Timestamp:   timestamp,
			EventType:   "UpdateClient",
			Name:        name,
			Email:       email,
		},
	}
}

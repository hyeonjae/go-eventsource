package events

import (
	"time"
)

type EnrollClient struct {
	Event
}

func NewEnrollClient(aggregateId string, timestamp time.Time, name, email string) *EnrollClient {
	return &EnrollClient{
		Event: Event{
			AggregateId: aggregateId,
			Timestamp:   timestamp,
			EventType:   "EnrollClient",
			Name:        name,
			Email:       email,
		},
	}
}

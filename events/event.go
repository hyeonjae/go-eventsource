package events

import (
	"bytes"
	"encoding/gob"
	"time"
)

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

func (e *Event) Encode() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(e); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (e *Event) Length() int {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(e); err != nil {
		return 0
	}
	return buf.Len()
}

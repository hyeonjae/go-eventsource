package eventstore

import (
	"context"

	eventsource "github.com/hyeonjae/go-eventsource"
	"github.com/hyeonjae/go-eventsource/events"
	"go.mongodb.org/mongo-driver/bson"
)

func (m MongoDB) FindAll(ctx context.Context, aggregateID string) ([]eventsource.Event, error) {
	collection := m.Database.Collection("events")

	cur, err := collection.Find(ctx, bson.M{
		"event.aggregate_id": aggregateID,
	})
	if err != nil {
		return []eventsource.Event{}, err
	}

	var evts []eventsource.Event
	for cur.Next(ctx) {
		row := struct {
			ID string `bson:"_id"`
			events.Event
		}{}
		cur.Decode(&row)
		switch row.Event.EventType {
		case "EnrollClient":
			evts = append(evts, &events.EnrollClient{Event: row.Event})
		case "UpdateClient":
			evts = append(evts, &events.UpdateClient{Event: row.Event})
		}
	}
	return evts, err
}

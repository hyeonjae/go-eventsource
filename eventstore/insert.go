package eventstore

import (
	"context"

	eventsource "github.com/hyeonjae/go-eventsource"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m MongoDB) Insert(ctx context.Context, events []eventsource.Event) ([]string, error) {
	collection := m.Database.Collection("events")

	var documents []interface{}
	for _, event := range events {
		documents = append(documents, event)
	}
	res, err := collection.InsertMany(ctx, documents)
	if err != nil {
		return []string{}, err
	}
	var ids []string
	for _, id := range res.InsertedIDs {
		ids = append(ids, id.(primitive.ObjectID).Hex())
	}
	return ids, err
}

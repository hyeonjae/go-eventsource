package clients

import (
	"context"
	"errors"

	eventsource "github.com/hyeonjae/go-eventsource"
	"github.com/hyeonjae/go-eventsource/eventstore"
)

type Service struct {
	eventStore *eventstore.MongoDB
}

func NewService(eventStore *eventstore.MongoDB) *Service {
	return &Service{
		eventStore: eventStore,
	}
}

func (s Service) Process(ctx context.Context, command ICommand) (string, error) {
	switch v := command.(type) {
	case EnrollCommand:
		client := v.Process()
		_, _ = s.eventStore.Insert(ctx, client.ListEvents())
		return client.GetID(), nil
	case UpdateCommand:
		client := v.Process()
		_, _ = s.eventStore.Insert(ctx, client.ListEvents())
		return client.GetID(), nil
	}
	return "", nil
}

func (s Service) Query(ctx context.Context, query IQuery) (eventsource.Aggregate, error) {
	switch v := query.(type) {
	case *GetQuery:
		client := v.Query()
		events, err := s.eventStore.FindAll(ctx, v.ID)
		if err != nil {
			return nil, err
		}
		for _, event := range events {
			client.Apply(event)
		}
		return client, nil
	}
	return nil, errors.New("invalid query type")
}

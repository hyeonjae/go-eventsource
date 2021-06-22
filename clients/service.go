package clients

import (
	"context"
	"errors"
	"fmt"

	eventsource "github.com/hyeonjae/go-eventsource"
	"github.com/hyeonjae/go-eventsource/eventbus"
	"github.com/hyeonjae/go-eventsource/events"
	"github.com/hyeonjae/go-eventsource/eventstore"
)

type Service struct {
	eventStore *eventstore.MongoDB
	eventBus   *eventbus.Kafka
}

func NewService(eventStore *eventstore.MongoDB, eventBus *eventbus.Kafka) *Service {
	return &Service{
		eventStore: eventStore,
		eventBus:   eventBus,
	}
}

func (s Service) Process(ctx context.Context, command ICommand) (string, error) {
	switch v := command.(type) {
	case EnrollCommand:
		client := v.Process()
		_, _ = s.eventStore.Insert(ctx, client.ListEvents())
		for _, event := range client.ListEvents() {
			switch v := event.(type) {
			case *events.EnrollClient:
				partition, offset, err := s.eventBus.Send("topic.clients", &v.Event)
				if err != nil {
					fmt.Printf("eventBus.Send err: %v\n", err)
					continue
				}
				fmt.Printf("partition: %d, offset: %d\n", partition, offset)
			}
		}
		return client.GetID(), nil
	case UpdateCommand:
		client := v.Process()
		_, _ = s.eventStore.Insert(ctx, client.ListEvents())
		for _, event := range client.ListEvents() {
			switch v := event.(type) {
			case *events.UpdateClient:
				partition, offset, err := s.eventBus.Send("topic.clients", &v.Event)
				if err != nil {
					fmt.Printf("eventBus.Send err: %v\n", err)
					continue
				}
				fmt.Printf("partition: %d, offset: %d\n", partition, offset)
			}
		}
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

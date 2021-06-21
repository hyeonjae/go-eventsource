package aggregate

import (
	"time"

	eventsource "github.com/hyeonjae/go-eventsource"
	"github.com/hyeonjae/go-eventsource/events"
)

type Client struct {
	ID     string
	Name   string
	Email  string
	Events []eventsource.Event
}

func EnrollClient(id string, timestamp time.Time, name, email string) *Client {
	client := Client{ID: id}
	enrollClient := events.NewEnrollClient(id, timestamp, name, email)
	client.Apply(enrollClient)
	client.Events = append(client.Events, enrollClient)
	return &client
}

func UpdateClient(id string, timestamp time.Time, name, email string) *Client {
	client := Client{ID: id}
	updateClient := events.NewUpdateClient(id, timestamp, name, email)
	client.Apply(updateClient)
	client.Events = append(client.Events, updateClient)
	return &client
}

func (c *Client) Apply(event eventsource.Event) {
	switch v := event.(type) {
	case *events.EnrollClient:
		c.Name = v.Name
		c.Email = v.Email
	case *events.UpdateClient:
		c.Name = v.Name
		c.Email = v.Email
	}
}

func (c Client) GetID() string {
	return c.ID
}

func (c Client) ListEvents() []eventsource.Event {
	return c.Events
}

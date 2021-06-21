package clients

import (
	"time"

	"github.com/google/uuid"
	eventsource "github.com/hyeonjae/go-eventsource"
	"github.com/hyeonjae/go-eventsource/aggregate"
)

type ICommand interface {
	Process() eventsource.Aggregate
}

type EnrollCommand struct {
	Name  string
	Email string
}

func (c EnrollCommand) Process() eventsource.Aggregate {
	id := uuid.Must(uuid.NewRandom()).String()
	timestamp := time.Now()
	return aggregate.EnrollClient(id, timestamp, c.Name, c.Email)
}

type UpdateCommand struct {
	ID    string
	Name  string
	Email string
}

func (c UpdateCommand) Process() eventsource.Aggregate {
	timestamp := time.Now()
	return aggregate.UpdateClient(c.ID, timestamp, c.Name, c.Email)
}

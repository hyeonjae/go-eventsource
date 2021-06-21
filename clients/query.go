package clients

import (
	eventsource "github.com/hyeonjae/go-eventsource"
	"github.com/hyeonjae/go-eventsource/aggregate"
)

type IQuery interface {
	Query() eventsource.Aggregate
}

type GetQuery struct {
	ID string
}

func (q GetQuery) Query() eventsource.Aggregate {
	return &aggregate.Client{
		ID: q.ID,
	}
}

package main

import (
	"fmt"
	"time"

	eventsource "github.com/hyeonjae/go-eventsource"
	"github.com/hyeonjae/go-eventsource/accounts"
	"github.com/hyeonjae/go-eventsource/clients"
	"github.com/hyeonjae/go-eventsource/datastore"
	"github.com/hyeonjae/go-eventsource/eventbus"
	"github.com/hyeonjae/go-eventsource/eventstore"
	"github.com/hyeonjae/go-eventsource/http"
	"github.com/hyeonjae/go-eventsource/kafka"
	"go.uber.org/fx"
)

var (
	startTimeout = 10 * time.Second
	stopTimeout  = 10 * time.Second
)

func main() {
	app := fx.New(
		fx.Provide(
			eventsource.NewConfig,
			datastore.New,
			eventbus.New,
			eventstore.New,
			clients.NewService,
			clients.NewResource,
			accounts.NewService,
			accounts.NewResource,
		),
		fx.Invoke(
			http.Start,
			kafka.Sink,
		),
		fx.StartTimeout(startTimeout),
		fx.StopTimeout(stopTimeout),
	)

	err := app.Err()
	if err != nil {
		fmt.Printf("run failed. error: %v\n", err)
		return
	}

	app.Run()
}

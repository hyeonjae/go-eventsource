package kafka

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"

	"github.com/hyeonjae/go-eventsource/datastore"
	"github.com/hyeonjae/go-eventsource/eventbus"
	"github.com/hyeonjae/go-eventsource/events"
	"github.com/pkg/errors"
)

func Sink(kafka *eventbus.Kafka, mysql *datastore.MySQL) error {
	msgch, errch, err := kafka.Receive("topic.clients")
	if err != nil {
		return errors.Wrapf(err, "fail to receive kafka")
	}

	go func() {
		for {
			select {
			case msg := <-msgch:
				fmt.Printf("msg: %v\n", msg)
				e := events.Event{}
				dec := gob.NewDecoder(bytes.NewReader(msg.Value))
				err := dec.Decode(&e)
				if err != nil {
					fmt.Printf("err: %v\n", err)
				}
				fmt.Printf("event: %v\n", e)
				ctx := context.Background()
				mysql.Upsert(ctx, &e)
			case err := <-errch:
				fmt.Printf("err: %v\n", err)
				break
			}
		}
	}()

	return nil
}

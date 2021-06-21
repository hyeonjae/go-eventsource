package eventstore

import (
	"context"
	"strings"

	eventsource "github.com/hyeonjae/go-eventsource"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Cli      *mongo.Client
	Database *mongo.Database
}

func New(cfg *eventsource.Config) (*MongoDB, error) {
	credential := options.Credential{
		Username:   cfg.Mongo.Username,
		Password:   cfg.Mongo.Password,
		AuthSource: cfg.Mongo.Database,
	}

	clientOpts := options.Client().ApplyURI(strings.Join(cfg.Mongo.Addresses, ",")).SetAuth(credential)
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		return nil, err
	}

	db := client.Database(cfg.Mongo.Database)

	return &MongoDB{
		client,
		db,
	}, nil
}

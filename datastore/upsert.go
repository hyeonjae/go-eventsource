package datastore

import (
	"context"

	"github.com/hyeonjae/go-eventsource/events"
)

func (m MySQL) Upsert(ctx context.Context, event *events.Event) error {
	_, err := m.DB.ExecContext(ctx, `INSERT INTO clients (id, name, email) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE name=?, email=?`, event.AggregateId, event.Name, event.Email, event.Name, event.Email)
	return err
}

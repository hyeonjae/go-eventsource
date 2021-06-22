package datastore

import (
	"fmt"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	eventsource "github.com/hyeonjae/go-eventsource"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type MySQL struct {
	*sqlx.DB
}

func New(cfg *eventsource.Config) (*MySQL, error) {
	db, err := connect(cfg.MySQL)
	return &MySQL{
		DB: db,
	}, err
}

func connect(cfg *eventsource.MySQL) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", buildDSN(cfg))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if err := db.Ping(); err != nil {
		return nil, errors.WithStack(err)
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(3600 * time.Second)

	return db, nil
}

func buildDSN(cfg *eventsource.MySQL) string {
	return fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=true",
		cfg.Username,
		url.PathEscape(cfg.Password),
		cfg.Address,
		cfg.Database)
}

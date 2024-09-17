package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"

	"checkmarx/internal/observer"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
	ErrDuplicateEmail = errors.New("duplicate email")
)

type Client struct {
	client *sql.DB
	observ *observer.Observer
}

func New(ctx context.Context, conf Configuration) (*Client, error) {
	db, err := sql.Open("pgx", conf.ConnectionURL)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(conf.MaxConnections)
	db.SetMaxIdleConns(conf.MaxIdleConnetion)
	db.SetConnMaxIdleTime(time.Duration(conf.MaxIdleConnetion))

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return &Client{client: db, observ: conf.Observer}, nil
}

func (c *Client) Close() error {
	if c.client != nil {
		return c.client.Close()
	}
	return nil
}

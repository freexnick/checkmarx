package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
	ErrDuplicateEmail = errors.New("duplicate email")
)

type Configuration struct {
	ConnectionURL    string
	MinConnections   int
	MaxConnections   int
	MaxIdleConnetion int
}

type Client struct {
	client *sql.DB
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

	return &Client{client: db}, nil
}

func (c *Client) Close() error {
	if c.client != nil {
		return c.client.Close()
	}
	return nil
}

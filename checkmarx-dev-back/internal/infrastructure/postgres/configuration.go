package postgres

import "checkmarx/internal/observer"

type Configuration struct {
	Observer         *observer.Observer
	ConnectionURL    string
	MinConnections   int
	MaxConnections   int
	MaxIdleConnetion int
}

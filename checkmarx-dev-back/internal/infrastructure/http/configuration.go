package server

import (
	"checkmarx/internal/observer"
	"net/http"
)

type Configuration struct {
	Observer     *observer.Observer
	Port         string
	Handler      http.Handler
	ReadTimeout  uint
	WriteTimeout uint
	IdleTimeout  uint
}

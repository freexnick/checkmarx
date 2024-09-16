package server

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type Configuration struct {
	Port         string
	Handler      http.Handler
	ReadTimeout  uint
	WriteTimeout uint
	IdleTimeout  uint
}

type Server struct {
	httpS *http.Server
}

func New(conf Configuration) (*Server, error) {
	if conf.Handler == nil || conf.Port == "" {
		return nil, errors.New("missing required options")
	}

	return &Server{
		httpS: &http.Server{
			Addr:         conf.Port,
			Handler:      conf.Handler,
			ReadTimeout:  time.Duration(conf.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(conf.WriteTimeout) * time.Second,
			IdleTimeout:  time.Duration(conf.IdleTimeout) * time.Second,
		},
	}, nil
}

func (s *Server) Start(ctx context.Context) error {
	return s.httpS.ListenAndServe()
}

func (s *Server) Close(ctx context.Context) error {
	err := s.httpS.Close()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

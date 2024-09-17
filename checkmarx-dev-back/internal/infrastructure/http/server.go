package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"checkmarx/internal/observer"
)

type Server struct {
	observ *observer.Observer
	httpS  *http.Server
}

func New(conf Configuration) (*Server, error) {
	if conf.Handler == nil || conf.Port == "" {
		return nil, errors.New("missing required options")
	}

	return &Server{
		observ: conf.Observer,
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
	s.observ.Info(ctx, fmt.Sprintf("HTTP server is listening on %s", s.httpS.Addr))

	return s.httpS.ListenAndServe()
}

func (s *Server) Close(ctx context.Context) error {
	err := s.httpS.Shutdown(ctx)
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

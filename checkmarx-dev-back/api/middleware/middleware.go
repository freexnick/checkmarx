package middleware

import (
	"checkmarx/internal/observer"
	"net/http"
	"time"
)

type MiddlewareHandler struct {
	Observer *observer.Observer
}

func New(c Configuration) *MiddlewareHandler {
	return &MiddlewareHandler{
		Observer: c.Observer,
	}
}

type KV = struct {
	Key   string
	Value any
}

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
func (m *MiddlewareHandler) Observe(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}
		start := time.Now()

		defer func() {
			m.Observer.Info(r.Context(), "Request completed",
				KV{Key: "status", Value: rw.statusCode},
				KV{Key: "duration", Value: time.Since(start)},
			)
		}()

		m.Observer.Info(r.Context(), "Request started",
			KV{Key: "method", Value: r.Method},
			KV{Key: "url", Value: r.URL.String()},
		)

		next.ServeHTTP(rw, r)

	})
}

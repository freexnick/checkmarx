package api

import (
	"github.com/go-chi/chi/v5"

	"checkmarx/api/handlers"
	"checkmarx/api/middleware"
)

type Configuration struct {
	Router            chi.Router
	MiddlewareHandler *middleware.MiddlewareHandler
	UserHandler       *handlers.UserHandler
	AuthHandler       *handlers.AuthHandler
	PostHandler       *handlers.PostHandler
	CommentHandler    *handlers.CommentHandler
	ApiVersion        uint
}

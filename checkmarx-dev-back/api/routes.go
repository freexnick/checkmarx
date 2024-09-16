package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"checkmarx/api/handlers"
)

type Configuration struct {
	Router         chi.Router
	UserHandler    *handlers.UserHandler
	AuthHandler    *handlers.AuthHandler
	PostHandler    *handlers.PostHandler
	CommentHandler *handlers.CommentHandler
	ApiVersion     uint
}

func New(conf Configuration) http.Handler {
	r := conf.Router
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Group(func(r chi.Router) {
		r.Route(fmt.Sprintf("/api/v%d", conf.ApiVersion), func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Route("/auth", func(r chi.Router) {
					r.With(conf.AuthHandler.Authenticate).Get("/", conf.AuthHandler.Get)
					r.Post("/signin", conf.AuthHandler.SignIn)
					r.Post("/signup", conf.AuthHandler.SignUp)
				})
			})
			r.Group(func(r chi.Router) {
				r.Use(conf.AuthHandler.Authenticate)
				r.Route("/posts", func(r chi.Router) {
					r.Get("/", conf.PostHandler.GetAllPosts)
					r.Post("/", conf.PostHandler.CreatePost)
					r.Get("/{id}", conf.PostHandler.GetPost)
					r.Put("/{id}", conf.PostHandler.UpdatePost)
					r.Delete("/{id}", conf.PostHandler.DeletePost)
				})
			})
			r.Group(func(r chi.Router) {
				r.Use(conf.AuthHandler.Authenticate)
				r.Route("/comments", func(r chi.Router) {
					r.Post("/", conf.CommentHandler.CreateComment)
					r.Put("/{id}", conf.CommentHandler.UpdateComment)
					r.Delete("/{id}", conf.CommentHandler.DeleteComment)
				})
			})
		})

	})

	return r
}

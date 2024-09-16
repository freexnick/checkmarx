package application

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"

	"checkmarx/api"
	"checkmarx/api/handlers"
	"checkmarx/internal/config"
	"checkmarx/internal/domain/service"
	httpS "checkmarx/internal/infrastructure/http"
	"checkmarx/internal/infrastructure/postgres"
)

type Application struct {
	postgresClient *postgres.Client

	userRepository    *postgres.UserRepository
	authRepository    *postgres.AuthRepository
	postRepository    *postgres.PostRepository
	commentRepository *postgres.CommentRepository

	userService    *service.UserService
	authService    *service.AuthService
	postService    *service.PostService
	commentService *service.CommentService

	userHandler    *handlers.UserHandler
	authHandler    *handlers.AuthHandler
	postHandler    *handlers.PostHandler
	commentHandler *handlers.CommentHandler

	handler    http.Handler
	httpServer *httpS.Server
}

func New(ctx context.Context) (*Application, error) {
	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	app := &Application{}

	if err := app.setRepositories(ctx, conf); err != nil {
		return nil, err
	}

	if err := app.setServices(); err != nil {
		return nil, err
	}

	if err := app.setRouteHandlers(); err != nil {
		return nil, err
	}

	if err := app.setRoutes(conf); err != nil {
		return nil, err
	}

	if err := app.setServers(conf); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *Application) setRepositories(ctx context.Context, conf *config.Configuration) error {
	db, err := postgres.New(ctx, postgres.Configuration{
		ConnectionURL:    conf.PostgresURL,
		MinConnections:   conf.PostgresMinConnections,
		MaxConnections:   conf.PostgresMaxConnections,
		MaxIdleConnetion: conf.PostgresMaxIdleTimeoutMinute,
	})
	if err != nil {
		return err
	}

	a.postgresClient = db

	a.userRepository = postgres.NewUserRepository(db)
	a.authRepository = postgres.NewAuthRepository(db)
	a.postRepository = postgres.NewPostRepository(db)
	a.commentRepository = postgres.NewCommentRepository(db)

	return nil
}

func (a *Application) setServices() error {
	a.userService = service.NewUserService(a.userRepository)
	a.authService = service.NewAuthService(a.authRepository)
	a.postService = service.NewPostService(a.postRepository)
	a.commentService = service.NewCommentService(a.commentRepository)

	return nil
}

func (a *Application) setRouteHandlers() error {
	a.userHandler = handlers.NewUserHandler(a.userService)
	a.authHandler = handlers.NewAuthHandler(a.authService)
	a.postHandler = handlers.NewPostHandler(a.postService)
	a.commentHandler = handlers.NewCommentHandler(a.commentService)

	return nil
}

func (a *Application) setRoutes(conf *config.Configuration) error {
	r := chi.NewRouter()

	routeConfig := api.Configuration{
		Router:         r,
		UserHandler:    a.userHandler,
		AuthHandler:    a.authHandler,
		PostHandler:    a.postHandler,
		CommentHandler: a.commentHandler,
		ApiVersion:     conf.ApiVersion,
	}

	a.handler = api.New(routeConfig)

	return nil
}

func (a *Application) setServers(conf *config.Configuration) error {
	httpS, err := httpS.New(httpS.Configuration{
		Port:         conf.HTTPServerAddress,
		Handler:      a.handler,
		ReadTimeout:  conf.ReadTimeoutSeconds,
		WriteTimeout: conf.WriteTimeoutSeconds,
		IdleTimeout:  conf.IdleTimeoutSeconds,
	})
	if err != nil {
		return err
	}

	a.httpServer = httpS
	return nil
}

func (a *Application) Start(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		if err := a.httpServer.Start(ctx); err != nil {
			return err
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}

func (a *Application) Close(ctx context.Context) error {
	if a.postgresClient != nil {
		if err := a.postgresClient.Close(); err != nil {
			return err
		}
	}

	if a.httpServer != nil {
		if err := a.httpServer.Close(ctx); err != nil {
			return err
		}
	}
	return nil
}

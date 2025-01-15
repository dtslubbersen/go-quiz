package api

import (
	"context"
	"errors"
	"fmt"
	swaggerDocs "github.com/dtslubbersen/go-quiz/docs"
	"github.com/dtslubbersen/go-quiz/internal/auth"
	"github.com/dtslubbersen/go-quiz/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

const version = "1.0.0"

type Application struct {
	authenticator auth.Authenticator
	cfg           apiCfg
	logger        *zap.SugaredLogger
	router        *chi.Mux
	storage       store.Storage
}

func NewApplication(ctx context.Context, logger *zap.SugaredLogger) *Application {
	cfg := apiCfg{
		apiUrl:        "localhost:8080",
		listenAddress: ":8080",
		authentication: authCfg{
			secret:      "notverysecret",
			expireAfter: time.Hour * 24 * 7, // 7 days
			iss:         "go-quiz",
		},
	}
	jwtAuthenticator := auth.NewJwtAuthenticator(
		cfg.authentication.secret,
		cfg.authentication.iss,
		cfg.authentication.iss,
	)

	seed := store.NewSeed()

	logger.Infoln("seed data extracted from disk")

	storage := store.NewStorage(seed)

	logger.Infoln("storage initialized")

	api := &Application{
		authenticator: jwtAuthenticator,
		cfg:           cfg,
		logger:        logger,
		storage:       storage,
	}

	api.setupRouter()

	return api
}

type apiCfg struct {
	apiUrl         string
	listenAddress  string
	authentication authCfg
}

type authCfg struct {
	secret      string
	expireAfter time.Duration
	iss         string
}

func (a *Application) setupRouter() {
	chiRouter := chi.NewRouter()

	chiRouter.Use(middleware.RequestID)
	chiRouter.Use(middleware.Logger)
	chiRouter.Use(middleware.Recoverer)
	chiRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{a.cfg.apiUrl},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           360,
	}))

	chiRouter.Use(middleware.Timeout(5 * time.Second))

	docsUrl := fmt.Sprintf("%s/swagger/doc.json", a.cfg.listenAddress)
	chiRouter.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(docsUrl)))

	chiRouter.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", a.healthCheckHandler)

		r.Route("/auth", func(r chi.Router) {
			r.Post("/token", a.createTokenHandler)
		})

		r.Route("/quizzes", func(r chi.Router) {
			r.Use(a.jwtTokenMiddleware)

			r.Get("/", a.getQuizzesHandler)

			r.Route("/{quizId}", func(r chi.Router) {
				r.Use(a.quizzesContextMiddleware)

				r.Get("/", a.getQuizByIdHandler)
				r.Post("/submit", a.postSubmitAnswersHandler)
				r.Get("/results", a.getQuizResultsHandler)
			})
		})
	})

	a.router = chiRouter
}

func (a *Application) Run() error {
	swaggerDocs.SwaggerInfo.Version = version
	swaggerDocs.SwaggerInfo.Host = a.cfg.apiUrl
	swaggerDocs.SwaggerInfo.BasePath = "/api/v1"

	server := &http.Server{
		Addr:         a.cfg.listenAddress,
		Handler:      a.router,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		a.logger.Infow("signal caught", "signal", s.String())

		shutdown <- server.Shutdown(ctx)
	}()

	a.logger.Infow("server started", "address", server.Addr)

	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdown
	if err != nil {
		return err
	}

	a.logger.Infow("server has stopped", "addr", a.cfg.listenAddress)

	return nil
}

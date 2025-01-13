package api

import (
	"context"
	"errors"
	"fmt"
	"go-quiz/internal/auth"
	"go-quiz/internal/store"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
	"net/http"
	"time"

	httpSwagger "github.com/swaggo/http-swagger/v2"
	swaggerDocs "go-quiz/docs"
)

type application struct {
	authenticator auth.Authenticator
	configuration appConfig
	logger        *zap.SugaredLogger
	store         store.Storage
}

type appConfig struct {
	apiUrl         string
	listenAddress  string
	authentication authConfig
}

type authConfig struct {
	secret      string
	expireAfter time.Duration
	iss         string
}

func (a *application) createRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{a.configuration.apiUrl},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           360,
	}))

	r.Use(middleware.Timeout(5 * time.Second))

	docsUrl := fmt.Sprintf("%s/swagger/doc.json", a.configuration.listenAddress)
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(docsUrl)))

	r.Route("/api/v1", func(r chi.Router) {
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
				r.Post("/submit", a.submitAnswersHandler)
				r.Get("/results", a.getQuizResultsHandler)
			})
		})
	})

	return r
}

func (a *application) run(mux http.Handler) error {
	swaggerDocs.SwaggerInfo.Version = version
	swaggerDocs.SwaggerInfo.Host = a.configuration.apiUrl
	swaggerDocs.SwaggerInfo.BasePath = "/api/v1"

	server := &http.Server{
		Addr:         a.configuration.listenAddress,
		Handler:      mux,
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

	a.logger.Infow("server has stopped", "addr", a.configuration.listenAddress)

	return nil
}

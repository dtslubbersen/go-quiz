package api

import (
	"expvar"
	"go-quiz/internal/auth"
	"go-quiz/internal/store"
	"go.uber.org/zap"
	"runtime"
	"time"
)

const version = "1.0.0"

//	@title			go-quiz
//	@version		1.0
//	@description	This is the API documentation for go-quiz, a simple Quiz API allowing users to obtain quizzes, answer the questions and see their results compared to other users.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Declan Lubbersen
//	@contact.url	https://github.com/dtslubbersen/go-quiz
//	@contact.email	dtslubbersen@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Description for what is this security definition being used

// @host		localhost:8080
// @BasePath	/api/v1
func Start() {
	cfg := appConfig{
		apiUrl:        "localhost:8080",
		listenAddress: ":8080",
		authentication: authConfig{
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

	logger := zap.Must(zap.NewDevelopment()).Sugar()
	defer logger.Sync()

	seed := store.NewSeed()

	logger.Infoln("seed data extracted from disk")

	store := store.NewStorage(seed)

	logger.Infoln("store initialized")

	app := &application{
		authenticator: jwtAuthenticator,
		configuration: cfg,
		logger:        logger,
		store:         store,
	}

	expvar.NewString("version").Set(version)
	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	mux := app.createRouter()

	logger.Fatal(app.run(mux))
}

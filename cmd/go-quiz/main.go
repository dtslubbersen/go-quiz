package main

import (
	"context"
	"github.com/dtslubbersen/go-quiz/internal/cmd"
	"os"
	"os/signal"
)

// @title						go-quiz
// @version					1.0
// @description				This is the API documentation for go-quiz, a simple Quiz API allowing users to obtain quizzes, answer the questions and see their results compared to other users.
// @termsOfService				http://swagger.io/terms/
//
// @contact.name				Declan Lubbersen
// @contact.url				https://github.com/dtslubbersen/go-quiz
// @contact.email				dtslubbersen@gmail.com
//
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//
// @scheme						http
// @host						localhost:8080
// @BasePath					/api/v1
//
// @securityDefinitions.apiKey	BearerAuth
// @type						apiKey
// @name						Authorization
// @in							header
// @description				Use a 'Bearer {token}' to authenticate your requests
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	ret := cmd.Execute(ctx)
	os.Exit(ret)
}

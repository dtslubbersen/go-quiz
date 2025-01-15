package cmd

import (
	"context"
	"github.com/dtslubbersen/go-quiz/internal/cli"
	openapi "github.com/dtslubbersen/go-quiz/pkg/client"
	"github.com/spf13/cobra"
	"log"
)

func CliCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cli",
		Short: "Runs the CLI",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := openapi.NewConfiguration()
			cfg.Servers = openapi.ServerConfigurations{
				{URL: "http://localhost:8080/api/v1"},
			}
			apiClient := openapi.NewAPIClient(cfg)

			quizContext := cli.NewQuizContext(apiClient, ctx)

			if err := quizContext.AuthenticateUser("demo@quiz.com", "password"); err != nil {
				log.Panicf("Authentication failed %v", err)
			}

			if err := quizContext.SelectQuiz(); err != nil {
				log.Panicf("Quiz selection failed %v", err)
			}

			if err := quizContext.CheckForExistingResults(); err != nil {
				return
			}

			if err := quizContext.AnswerQuestions(); err != nil {
				log.Panicf("Failed to process answers %v", err)
			}

			if err := quizContext.SubmitAnswers(); err != nil {
				log.Panicf("Failed to submit answers %v", err)
			}

			if err := quizContext.DisplayResults(); err != nil {
				log.Panicf("Failed to display results %v", err)
			}
		},
	}

	return cmd
}

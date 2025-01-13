package cli

import (
	"context"
	"github.com/spf13/cobra"
	openapi "go-quiz/pkg/client"
	"log"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a quiz",
	Run:   handleStartCmd,
}

func handleStartCmd(cmd *cobra.Command, args []string) {
	cfg := openapi.NewConfiguration()
	cfg.Servers = openapi.ServerConfigurations{
		{URL: "http://localhost:8080/api/v1"},
	}
	apiClient := openapi.NewAPIClient(cfg)
	ctx := context.Background()

	quizContext := NewQuizContext(apiClient, ctx)

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
}

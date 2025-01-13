package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const quizCtxKey = "quiz"

var rootCmd = &cobra.Command{
	Use:   "go-quiz-cli",
	Short: "go-quiz-cli is a simple quiz CLI that interacts with a quiz backend",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to go-quiz-cli")
	},
}

type appConfig struct {
	apiUrl string
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

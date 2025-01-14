package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

func Execute(ctx context.Context) int {
	rootCmd := &cobra.Command{
		Use:   "go-quiz",
		Short: "go-quiz is a simple quiz game",
	}

	rootCmd.AddCommand(ApiCmd(ctx))
	rootCmd.AddCommand(CliCmd(ctx))

	if err := rootCmd.Execute(); err != nil {
		return 1
	}

	return 0
}

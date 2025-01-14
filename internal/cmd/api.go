package cmd

import (
	"context"
	"github.com/dtslubbersen/go-quiz/internal/api"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func ApiCmd(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Args:  cobra.NoArgs,
		Short: "Runs the RESTful API",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := zap.Must(zap.NewDevelopment()).Sugar()
			defer func() { _ = logger.Sync() }()

			application := api.NewApplication(ctx, logger)
			return application.Run()
		},
	}

	return cmd
}

package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"platformctl/internal/action/service"
	"platformctl/internal/cfg"
	"platformctl/internal/minikube"
)

var serviceDebugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Debug a service",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutMediumOperation())
		defer cancel()

		if _, err := minikube.IsRunning(ctx); err != nil {
			return err
		}

		return service.Debug(ctx)
	},
}

func init() {
	serviceCmd.AddCommand(serviceDebugCmd)
}
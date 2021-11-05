package cmd

import (
	"context"

	"platformctl/internal/cfg"
	"platformctl/internal/minikube"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop minikube",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), cfg.TimeoutMediumOperation())
		defer cancel()

		return minikube.Stop(ctx)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
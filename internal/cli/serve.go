package cli

import (
	"github.com/arielril/hexgo/internal/handler"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start a HTTP server",
		RunE: func(cmd *cobra.Command, args []string) error {
			httpHandler := handler.NewHttpServer(ctx)

			return httpHandler.Serve()
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

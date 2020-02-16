package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	echoCmd = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}
)

func init() {
	rootCmd.AddCommand(echoCmd)
}

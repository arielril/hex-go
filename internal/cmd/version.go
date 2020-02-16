package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the current version of the cli",
	Long:  `All software has versions. This is CLI's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Golang CLI Template %s\n", viper.GetString("version"))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

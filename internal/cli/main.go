package cli

import (
	"fmt"
	"os"

	"github.com/arielril/hexgo/internal/handler"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "cli",
		Short: "Golang CLI Template project",
	}
	ctx *handler.HandlerContext
)

// Executes the command line interface
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is ./configs/.config.yaml)")
}

func initContext() {
	ctx = handler.NewContext()
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./configs")
		viper.SetConfigName(".config")
	}

	viper.SetConfigType("yaml")

	viper.SetEnvPrefix("cli")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to load config file:", viper.ConfigFileUsed())
	}

	initContext()
}

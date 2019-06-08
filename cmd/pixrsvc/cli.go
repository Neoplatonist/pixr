package main

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var logger = log.New(io.Writer(os.Stdout), "cli:", log.Lshortfile)

func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")  // name of config file (without extension)
	viper.AddConfigPath("configs") // path to look for the config file in
	if err := viper.ReadInConfig(); err != nil {
		logger.Println("could not read config file")
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// Execute starts the root cli command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func main() {
	Execute()
}

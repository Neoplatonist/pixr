package main

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var logger = log.New(io.Writer(os.Stdout), "cli:", log.Lshortfile)

func initConfig() {

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

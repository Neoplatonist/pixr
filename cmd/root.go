package main

import "github.com/spf13/cobra"

var (
	rootOptions struct{}

	rootCmd = &cobra.Command{
		Use:   "pixr",
		Short: "Pixr is a simplistic image sharing server.",
		Long:  "Pixr is a simplistic image sharing server.",
		RunE:  rootCmdFunc,
	}
)

func init() {
	// root command cli flags go here
}

func rootCmdFunc(cmd *cobra.Command, args []string) error {
	return nil
}

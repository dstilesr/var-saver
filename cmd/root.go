/*
Copyright © 2025 NAME HERE
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "var-saver [command] [options]",
	Short: "VarSaver helps you manage variables and values in a local environment",
	Long: `This is a solution to help save and retrieve variables to save you
the pain of looking through 20 env and config files whenever you need a
specific API URL you used once a couple weeks ago!
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

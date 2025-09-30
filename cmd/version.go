/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	confighandle "var-saver/configHandle"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the app version",
	Long:  `Prints the application version and metadata`,
	Run: func(cmd *cobra.Command, args []string) {
		sv := confighandle.ReadConfig()
		confighandle.PrintItem(sv.Meta)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

/*
Copyright © 2025 NAME HERE
*/
package cmd

import (
	"fmt"
	"strings"
	cfgh "var-saver/configHandle"

	"github.com/spf13/cobra"
)

// RunGetValue Runs the command to get a variable's value
func RunGetValue(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	env, _ := cmd.Flags().GetString("environment")
	proj, _ := cmd.Flags().GetString("project")

	sv := cfgh.ReadConfig()

	pNorm := strings.ToLower(proj)
	p, ok := sv.Projects[pNorm]
	if !ok {
		return fmt.Errorf("project '%s' not found", proj)
	}
	v, err := p.GetVariable(name, env)
	if err != nil {
		return err
	}
	fmt.Println(v.Value)
	return nil
}

// getValueCmd represents the getValue command
var getValueCmd = &cobra.Command{
	Use:   "getValue --name <var name> --project <project name> --environment <value environment>",
	Short: "Get the value of a variable",
	Long:  "Get the value of a variable given its name, project, and environment",
	RunE:  RunGetValue,
}

func init() {
	rootCmd.AddCommand(getValueCmd)

	getValueCmd.Flags().StringP("name", "n", "", "Name of the variable")
	getValueCmd.Flags().StringP("environment", "e", "", "Name of the environment")
	getValueCmd.Flags().StringP("project", "p", "", "Name of the project")

	getValueCmd.MarkFlagRequired("name")
	getValueCmd.MarkFlagRequired("environment")
	getValueCmd.MarkFlagRequired("project")
}

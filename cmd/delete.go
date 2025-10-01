/*
Copyright © 2025 NAME HERE
*/
package cmd

import (
	"errors"
	"strings"
	confighandle "var-saver/configHandle"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete {variable | project} [options]",
	Short: "Delete an item from the stored config",
	Long: `Delete a variable or a project from the stored config

Use examples:

# Delete a project and all its variables
delete project --name my-project

# Delete a specific variable
delete variable --project my-project --name my-var --environment dev
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("specify the kind of item to delete")
	},
}

var deleteProject = &cobra.Command{
	Use:   "project --name <project-name>",
	Short: "Delete a project and all its variables",
	Long: `Delete a project from the stored config

Use example:

delete project --name my-project
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		project, _ := cmd.Flags().GetString("project")
		project = strings.ToLower(strings.TrimSpace(project))

		sv := confighandle.ReadConfig()
		defer sv.SaveCfg()

		return sv.RemoveProject(project)
	},
}

var deleteVariable = &cobra.Command{
	Use:   "variable --name <var-name> --environment <env-name> --project <project-name>",
	Short: "Delete a project and all its variables",
	Long: `Delete a project from the stored config

Use example:

delete variable --name my-var --project my-project --environment dev
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		project, _ := cmd.Flags().GetString("project")
		name, _ := cmd.Flags().GetString("name")
		env, _ := cmd.Flags().GetString("environment")

		project = strings.ToLower(strings.TrimSpace(project))
		name = strings.ToLower(strings.TrimSpace(name))
		env = strings.ToLower(strings.TrimSpace(env))

		sv := confighandle.ReadConfig()
		defer sv.SaveCfg()

		return sv.RemoveVariable(project, name, env)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Delete project setup
	deleteCmd.AddCommand(deleteProject)
	deleteProject.Flags().StringP("name", "n", "", "Name of project to delete")
	deleteProject.MarkFlagRequired("name")

	// Delete variable setup
	deleteCmd.AddCommand(deleteVariable)
	deleteVariable.Flags().StringP("name", "n", "", "Name of variable to delete")
	deleteVariable.Flags().StringP("project", "p", "", "Name of project from which to delete variable")
	deleteVariable.Flags().StringP("environment", "e", "", "Environment of variable to delete")

	deleteVariable.MarkFlagRequired("name")
	deleteVariable.MarkFlagRequired("project")
	deleteVariable.MarkFlagRequired("environment")

}

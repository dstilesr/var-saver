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

// ValidateCreate Validates the inpút flags / args to the create command
func ValidateCreate(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	project, _ := cmd.Flags().GetString("project")
	env, _ := cmd.Flags().GetString("environment")
	val, _ := cmd.Flags().GetString("value")

	// Check inputs
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("ERROR: You must provide a nonempty variable name")
	}

	if strings.TrimSpace(project) == "" {
		return fmt.Errorf("ERROR: Project name cannot be empty")
	}

	if strings.TrimSpace(env) == "" {
		return fmt.Errorf("ERROR: Environment name cannot be empty")
	}

	if strings.TrimSpace(val) == "" {
		return fmt.Errorf("ERROR: You must provide a nonempty value")
	}
	return nil
}

// RunCreate Runs the command to add a new variable
func RunCreate(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	project, _ := cmd.Flags().GetString("project")
	env, _ := cmd.Flags().GetString("environment")
	val, _ := cmd.Flags().GetString("value")
	ow, _ := cmd.Flags().GetBool("overwrite")

	project = strings.TrimSpace(project)
	name = strings.TrimSpace(name)
	val = strings.TrimSpace(val)
	env = strings.TrimSpace(env)

	sv := cfgh.ReadConfig()
	defer sv.SaveCfg()

	err := sv.AddVariable(project, name, env, val, ow)
	if err != nil {
		return err
	}

	fmt.Printf(
		"Variable '%s' for project '%s' (%s) added successfully!\n",
		name,
		project,
		env,
	)
	return nil
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create --name <variable name> --value <variable value> [--project <project name>] [--environment <env name>] [--overwrite]",
	Short: "Create a new variable.",
	Long: `Create a new variable for later reference.
To create a variable you must provide the variable name and value. You can
optionally provide a project name and environment, but these will default to
"common" and "default" respectively. To overwrite an existing variable, use
the '--overwrite' flag.

Usage examples:

var-saver create --name "some-api-url" --value "https://my-api.com"

var-saver create --name "some-api-url" --value "https://my-api.com" --overwrite

var-saver create --name "some-api-url" --value "https://my-api.com" --project "my-project" --environment "dev"
`,
	RunE:    RunCreate,
	PreRunE: ValidateCreate,
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Flags + settings
	createCmd.Flags().StringP("name", "n", "", "Name of variable to add")
	createCmd.Flags().String("value", "", "Value of the variable")

	createCmd.Flags().StringP("project", "p", "common", "Name of project related to variable")
	createCmd.Flags().StringP(
		"environment",
		"e",
		"default",
		"Name of environment associated to var to add",
	)
	createCmd.Flags().Bool(
		"overwrite",
		false,
		"Whether to overwrite the variable if it exists",
	)

	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("value")
}

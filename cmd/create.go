/*
Copyright © 2025 NAME HERE
*/
package cmd

import (
	"errors"
	"fmt"
	"strings"
	cfgh "var-saver/configHandle"

	"github.com/spf13/cobra"
)

// ValidateCreateVar Validates the inpút flags / args to the create command
func ValidateCreateVar(cmd *cobra.Command, args []string) error {
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

// RunCreateVar Runs the command to add a new variable
func RunCreateVar(cmd *cobra.Command, args []string) error {
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

// ValidateCreatePrj Validates the inpút flags / args to the create command
func ValidateCreatePrj(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")

	// Check inputs
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("ERROR: You must provide a nonempty project name")
	}
	return nil
}

// RunCreatePrj Runs the command to add a new project
func RunCreatePrj(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	description, _ := cmd.Flags().GetString("description")

	description = strings.TrimSpace(description)
	name = strings.TrimSpace(name)

	sv := cfgh.ReadConfig()
	defer sv.SaveCfg()

	err := sv.AddProject(name, description)
	if err != nil {
		return err
	}

	fmt.Printf("Project '%s' added successfully!\n", name)
	return nil
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create {variable|project} [options]",
	Short: "Create a new variable or project.",
	Long:  `Create a new variable or project for later reference.`,
	RunE: func(c *cobra.Command, args []string) error {
		return errors.New("you must specify whether you want to create a variable or project")
	},
}

// createVar represents the command to create a new variable
var createVar = &cobra.Command{
	Use:   "variable --name <variable name> --value <variable value> [--project <project name>] [--environment <env name>] [--overwrite]",
	Short: "Create a new variable.",
	Long: `Create a new variable for later reference.
To create a variable you must provide the variable name and value. You can
optionally provide a project name and environment, but these will default to
"common" and "default" respectively. To overwrite an existing variable, use
the '--overwrite' flag.

Usage examples:

var-saver create variable --name "some-api-url" --value "https://my-api.com"

var-saver create variable --name "some-api-url" --value "https://my-api.com" --overwrite

var-saver create variable --name "some-api-url" --value "https://my-api.com" --project "my-project" --environment "dev"
`,
	RunE:    RunCreateVar,
	PreRunE: ValidateCreateVar,
}

// createPrj represents the command to create a project
var createPrj = &cobra.Command{
	Use:   "project --name '<project name>' [--description '<description for the project>']",
	Short: "Create a new project",
	Long: `Create a new project to associate variables.
You may provide a description for the project as well as a name.

Usage Example:
var-saver create project --name "my-proj" --description "This is a really cool project"
`,
	RunE:    RunCreatePrj,
	PreRunE: ValidateCreatePrj,
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Create Variable command
	createCmd.AddCommand(createVar)
	createVar.Flags().StringP("name", "n", "", "Name of variable to add")
	createVar.Flags().String("value", "", "Value of the variable")

	createVar.Flags().StringP("project", "p", "common", "Name of project related to variable")
	createVar.Flags().StringP(
		"environment",
		"e",
		"default",
		"Name of environment associated to var to add",
	)
	createVar.Flags().Bool(
		"overwrite",
		false,
		"Whether to overwrite the variable if it exists",
	)

	createVar.MarkFlagRequired("name")
	createVar.MarkFlagRequired("value")

	// Create Project command
	createCmd.AddCommand(createPrj)
	createPrj.Flags().StringP("name", "n", "", "Name of project to add")
	createPrj.Flags().StringP("description", "d", "", "Description for the project")

	createPrj.MarkFlagRequired("name")
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strings"
	confighandle "var-saver/configHandle"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list {projects | variables} [options]",
	Short: "List items of the given type",
	Long: `List projects or variables

Use examples:
list projects

# List variables in a specific project
list variables --project my-project
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("please enter with a subcommand specifying the objects to list!")
	},
}

var listProjects = &cobra.Command{
	Use:   "projects [--detail]",
	Short: "List projects",
	Long: `List stored projects

Use examples:

projects

# Print projects with their associated variables as well
projects --detail
`,
	Run: func(cmd *cobra.Command, args []string) {
		sv := confighandle.ReadConfig()
		detail, _ := cmd.Flags().GetBool("detail")
		fmt.Println("Projects:")
		for _, p := range sv.Projects {
			if detail {
				fmt.Printf("- '%s' (%d variables)\n", p.Name, len(p.Variables))
			} else {
				confighandle.PrintItem(p)
			}
		}
	},
}

var listVariables = &cobra.Command{
	Use:   "variables --project <project-name> [--environment <env name>]",
	Short: "List variables",
	Long: `List the variables for a given project

Use example:

variables --project my-project

variables -p my-project --environment dev
`,
	Run: func(cmd *cobra.Command, args []string) {
		sv := confighandle.ReadConfig()
		project, _ := cmd.Flags().GetString("project")
		env, _ := cmd.Flags().GetString("environment")

		project = strings.ToLower(strings.TrimSpace(project))
		env = strings.ToLower(strings.TrimSpace(env))

		p, ok := sv.Projects[project]
		if !ok || len(p.Variables) == 0 {
			fmt.Printf("No variables found for project '%s'", project)
			return
		}
		for _, v := range p.Variables {
			if env == "" || env == v.Environment {
				confighandle.PrintItem(v)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// List Projects setup
	listCmd.AddCommand(listProjects)
	listProjects.Flags().BoolP(
		"detail",
		"d",
		false,
		"Whether to give more project details including variables",
	)

	// List vars setup
	listCmd.AddCommand(listVariables)
	listVariables.Flags().StringP("project", "p", "", "Project to list variables")
	listVariables.Flags().StringP("environment", "e", "", "List only variables in this environment")
	listVariables.MarkFlagRequired("project")
}

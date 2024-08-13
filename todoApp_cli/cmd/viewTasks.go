/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"flag"
	"github.com/Ashbeeson7943/GO_Projects/todoApp_cli/task"
	"github.com/spf13/cobra"
)

// viewTasksCmd represents the viewTasks command
var viewTasksCmd = &cobra.Command{
	Use:   "viewTasks",
	Short: "View tasks on your task list",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		task.ViewTasks(at)
	},
}

var at bool

func init() {
	rootCmd.AddCommand(viewTasksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewTasksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewTasksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	viewTasksCmd.Flags().BoolVarP(&at, "all", "a", false, "Shows all tasks including those completed")
	flag.Parse()
}

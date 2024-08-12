/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Ashbeeson7943/GO_Projects/todoApp_cli/task"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// addTaskCmd represents the addTask command
var addTaskCmd = &cobra.Command{
	Use:   "addTask",
	Short: "Add a task to your todo list",
	Long:  `Add a task to your todo list`,
	Run: func(cmd *cobra.Command, args []string) {
		addTask(args)
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addTask(ts []string) {
	s := strings.Join(ts[:], " ")
	t := task.Task{
		ID:               0,
		TASK_TITLE:       s,
		TASK_DETAIL:      "",
		CREATED_TIME:     time.Now(),
		IS_COMPLETED:     false,
		COMPLETED_REASON: "",
		COMPLETED_TIME:   time.Time{},
	}
	task.SaveTask(t)
	fmt.Printf("%+v", t)
}

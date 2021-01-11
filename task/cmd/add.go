package cmd

import (
	"fmt"
	"strings"
	"../db"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.AddCommand(reviewCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			return
		}
		fmt.Printf("Added \"%s\" to your task.\n", task)
	},
}

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Add a new task to your TODO list",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is a fake \"review\" command")
	},
}
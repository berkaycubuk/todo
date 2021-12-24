/*
Copyright © 2021 Berkay Çubuk <berkay@berkaycubuk.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task`,
	Run: func(cmd *cobra.Command, args []string) {
		addTask(args[0])
	},
}

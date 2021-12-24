/*
Copyright © 2021 Berkay Çubuk <berkay@berkaycubuk.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information of Todo",
	Long:  `Long version information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Todo v0.1")
	},
}

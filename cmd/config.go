/*
Copyright © 2023 TJ Gillis <tj@doublebacklabs.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "create a new config",
	Long:  `If no subcommand is passed, it will default to the create subcommand`,
	Run: func(cmd *cobra.Command, args []string) {
		createCmd.Execute()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

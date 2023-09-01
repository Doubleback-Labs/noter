/*
Copyright © 2023 TJ Gillis <tj@doublebacklabs.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create new config",
	Long:  `create a central config that will ensure your notes are all gathered in the same location`,
	Run: func(cmd *cobra.Command, args []string) {
		Save()
	},
}

func init() {
	configCmd.AddCommand(createCmd)
}

func Save() {
	viper.WriteConfig()
}

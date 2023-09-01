/*
Copyright © 2023 TJ Gillis <tj@doublebacklabs.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "create a new config",
	Long:  `If no subcommand is passed, it will default to the create subcommand`,
	Run: func(cmd *cobra.Command, args []string) {
		Save()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	viper.BindPFlag("contentDir", rootCmd.PersistentFlags().Lookup("contentDir"))
	viper.BindPFlag("editor", configCmd.PersistentFlags().Lookup("editor"))

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.SetDefault("contentDir", fmt.Sprintf("%s/.noter/notes", home))
	viper.SetDefault("editor", "code")
}

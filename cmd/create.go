/*
Copyright © 2023 TJ Gillis <tj@doublebacklabs.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create new config",
	Long:  `create a central config that will ensure your notes are all gathered in the same location`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Executing `create` command")

		viper.SafeWriteConfig()
	},
}

func init() {
	configCmd.AddCommand(createCmd)

	viper.BindPFlag("contentDir", rootCmd.PersistentFlags().Lookup("contentDir"))
	viper.BindPFlag("editor", configCmd.PersistentFlags().Lookup("editor"))

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.SetDefault("contentDir", fmt.Sprintf("%s/.noter/notes", home))
	viper.SetDefault("editor", "code")

}

/*
Copyright © 2023 TJ Gillis <tj@doublebacklabs.com>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list notes",
	Long:  `lists all notes in your note repo`,
	Run: func(cmd *cobra.Command, args []string) {
		contentPath := viper.GetString("contentDir")
		fmt.Printf("notes:\n")
		err := filepath.Walk(contentPath, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
				return err
			}
			fmt.Printf("   - %q\n", filepath.Base(path))
			return nil
		})

		if err != nil {
			log.Fatal().Msg(err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

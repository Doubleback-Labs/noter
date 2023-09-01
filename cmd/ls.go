/*
Copyright © 2023 TJ Gillis <tj@doublebacklabs.com>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		contentPath := viper.GetString("contentDir")
		fmt.Printf("notes:\n")
		err := filepath.Walk(contentPath, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
				return err
			}
			fmt.Printf("	- %q\n", filepath.Base(path))
			return nil
		})
		if err != nil {
			fmt.Printf("error walking the path %q: %v\n", contentPath, err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

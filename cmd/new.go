/*
Copyright © 2023 TJ Gillis <tj@doublebacklabs.com>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var contentName string
var hugoPost bool

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "new note",
	Long: `Will open your configred editor with a filepath set to your central note repo.

If no name is provided, the default is to use the user date. 

If the hugo flag is passed, it will will assume that your central note repo is the base of your hugo site and try
create the while in the posts folder
`,
	Run: func(cmd *cobra.Command, args []string) {
		NewPost()
	},
}

func NewPost() {
	contentPath := viper.GetString("contentDir")
	contentName := viper.GetString("contentName")
	editor := viper.GetString("editor")

	if hugoPost {
		contentPath = newHugoContent(contentPath, contentName)
	}

	editorCommend := exec.Command(editor, fmt.Sprintf("%s.md", contentName))
	editorCommend.Dir = contentPath
	err := editorCommend.Run()
	if err != nil {
		fmt.Printf("Err %v", err)
	}
}

func newHugoContent(contentPath string, name string) string {
	hugoCmd := exec.Command("hugo", "new", "content", fmt.Sprintf("posts/%v.md", name))
	hugoCmd.Dir = contentPath
	if err := hugoCmd.Run(); err != nil {
		log.Fatal().Msgf("hugo err %v", err)
	}

	return fmt.Sprintf("%s/content/posts", contentPath)
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.PersistentFlags().StringVarP(&contentName, "contentName", "c", "", "content directory (default is $HOME/.noter/notes)")
	newCmd.PersistentFlags().BoolVarP(&hugoPost, "hugoPost", "p", false, "content directory (default is $HOME/.noter/notes)")
	viper.BindPFlag("contentName", newCmd.PersistentFlags().Lookup("contentName"))
	viper.BindPFlag("hugoPost", newCmd.PersistentFlags().Lookup("hugoPost"))

	viper.SetDefault("contentName", time.DateOnly)
	viper.SetDefault("hugoPost", false)
}

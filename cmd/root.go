/*
Copyright © 2023 TJ Gillis <tj@doublebacklabs.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var contentDir string
var editor string
var contentName string
var hugoPost bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "noter",
	Short: "noter helps make note taking easier.",
	Long: `noter is a simple app to open a GUI (atm) editor of your choice assuming it has a 'app filename' command.
Files are stored in a central repo of your choosing.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		NewPost()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.noter/cfg.yaml)")
	rootCmd.PersistentFlags().StringVar(&contentDir, "contentDir", "", "content directory (default is $HOME/.noter/notes)")
	rootCmd.PersistentFlags().StringVar(&editor, "editor", "", "editor that can be opened like 'app filename'")
	rootCmd.PersistentFlags().StringVarP(&contentName, "contentName", "c", time.Now().Format(time.DateOnly), "content directory (default is $HOME/.noter/notes)")
	rootCmd.PersistentFlags().BoolVarP(&hugoPost, "hugoPost", "p", false, "content directory (default is $HOME/.noter/notes)")
	viper.BindPFlag("contentName", rootCmd.PersistentFlags().Lookup("contentName"))
	viper.BindPFlag("hugoPost", rootCmd.PersistentFlags().Lookup("hugoPost"))
	viper.BindPFlag("contentDir", rootCmd.PersistentFlags().Lookup("contentDir"))
	viper.BindPFlag("editor", rootCmd.PersistentFlags().Lookup("editor"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".noter" (without extension).
		viper.AddConfigPath(fmt.Sprintf("%s/.noter", home))
		viper.SetConfigType("yaml")
		viper.SetConfigName("cfg")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)

		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		create(fmt.Sprintf("%s/.noter/cfg.yaml", home))

		if err := os.MkdirAll(fmt.Sprintf("%s/.noter/notes", home), 0770); err != nil {
			fmt.Println(err)
		}
	}
}

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func NewPost() {
	contentPath := viper.GetString("contentDir")
	contentName := viper.GetString("contentName")
	editor := viper.GetString("editor")

	if hugoPost {
		contentPath = newHugoContent(contentPath, contentName)
	}

	contentName = fmt.Sprintf("%s.md", contentName)

	f, err := os.OpenFile(fmt.Sprintf("%s/%s", contentPath, contentName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Debug().Msg(err.Error())
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("\n## %s\n", time.Now().Format(time.TimeOnly))); err != nil {
		log.Debug().Msg(err.Error())
	}

	editorCommend := exec.Command(editor, contentName)
	editorCommend.Dir = contentPath
	err = editorCommend.Run()
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

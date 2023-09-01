/*
Copyright © 2023 TJ Gillis <tj@doublebacklabs.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var contentDir string
var editor string
var debug bool

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
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "config file (default is $HOME/.noter/cfg.yaml)")
	rootCmd.PersistentFlags().StringVar(&contentDir, "contentDir", "", "content directory (default is $HOME/.noter/notes)")
	rootCmd.PersistentFlags().StringVar(&editor, "editor", "", "editor that can be opened like 'app filename'")

	viper.BindPFlag("contentDir", rootCmd.PersistentFlags().Lookup("contentDir"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	viper.SetDefault("debug", false)

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

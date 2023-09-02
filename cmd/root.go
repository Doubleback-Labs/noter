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

const typedNoterNoteRepo string = "noteRepo"
const typedNoterName string = "name"
const typedNoterEditor string = "editor"
const typedNoterIsHugo string = "isHugoPost"

type NoterFileData struct {
	Path        string
	Name        string
	Editor      string
	IsHugo      bool
	DefaultName bool
}

var cfgFile string
var noteRepo string
var editor string
var name string
var isHugoPost bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "noter",
	Short: "noter helps make note taking easier.",
	Long: `noter is a simple app to open an editor of your choice assuming it has a 'app filename' command.
	
	eg: code hello.md or micro hello.md

Files are stored in a central repo of your choosing.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		Post(GetFilePath())
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
	home, _ := os.UserHomeDir()
	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", fmt.Sprintf("%s/.noter/cfg.yaml", home), "config file (default is $HOME/.noter/cfg.yaml)")
	rootCmd.Flags().StringVarP(&noteRepo, typedNoterNoteRepo, "r", fmt.Sprintf("%s/.noter/notes", home), "content directory (default is $HOME/.noter/notes)")
	rootCmd.Flags().StringVarP(&editor, typedNoterEditor, "e", "nano", "editor that can be opened like 'app filename'")
	// Default value for name left intentinally empty to determine if using default name or not
	rootCmd.Flags().StringVarP(&name, typedNoterName, "n", "", "defaults to DateOnly name (yyyy-mm-dd)")
	rootCmd.Flags().BoolVarP(&isHugoPost, typedNoterIsHugo, "p", false, "If true will use hugo-cli to create and open post")

	//viper.BindPFlag(typedNoterName, rootCmd.Flags().Lookup(typedNoterName))
	//viper.BindPFlag(typedNoterIsHugo, rootCmd.Flags().Lookup(typedNoterIsHugo))
	viper.BindPFlag(typedNoterNoteRepo, rootCmd.Flags().Lookup(typedNoterNoteRepo))
	viper.BindPFlag(typedNoterEditor, rootCmd.Flags().Lookup(typedNoterEditor))

	cobra.OnInitialize(initConfig)

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

		createConfig(fmt.Sprintf("%s/.noter/cfg.yaml", home))
		viper.WriteConfig()

		if err := os.MkdirAll(fmt.Sprintf("%s/.noter/notes", home), 0770); err != nil {
			fmt.Println(err)
		}
	}
}

func createConfig(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func GetFilePath() NoterFileData {
	noteRepo := viper.GetString(typedNoterNoteRepo)
	editor := viper.GetString(typedNoterEditor)
	isHugo := viper.GetBool(typedNoterIsHugo)
	defaultName := false

	if name == "" {
		name = time.Now().Format(time.DateOnly)
		defaultName = true
	}

	if isHugo {
		noteRepo = newHugoContent(noteRepo, name)
	}

	return NoterFileData{
		Path:        noteRepo,
		Name:        name,
		Editor:      editor,
		IsHugo:      isHugo,
		DefaultName: defaultName,
	}
}

func newHugoContent(path string, name string) string {
	cmd := exec.Command("hugo", "new", "content", fmt.Sprintf("posts/%v.md", name))
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		log.Fatal().Msgf("hugo err %v", err)
	}

	return fmt.Sprintf("%s/content/posts/%s.md", path, name)
}

func Post(d NoterFileData) {
	f, err := os.OpenFile(fmt.Sprintf("%s/%s", d.Path, d.Name), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Debug().Msg(err.Error())
	}

	defer f.Close()

	if d.DefaultName {
		if _, err := f.WriteString(fmt.Sprintf("\n## %s\n", time.Now().Format(time.TimeOnly))); err != nil {
			log.Debug().Msg(err.Error())
		}
	}

	cmd := exec.Command(d.Editor, d.Name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Dir = d.Path
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Err %v", err)
	}
}

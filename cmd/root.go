package cmd

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-projects-launcher [pattern]",
	Short: "A tool that allows for eaching switching of projects",
	Long:  `Through the config file, the tool knows where to look for projects, apply any filters and determine possible IDE to launch the project for`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: MainWorkload,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		SetupLogger(cmd, args)
		log.Info("==== project switcher ====")
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-projects-launcher.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("config", "c", false, "Edit the config")
	rootCmd.Flags().BoolP("setup", "s", false, "Go through basic setup")
	rootCmd.Flags().String("ide", "", "The specific ide to use")
}

func MainWorkload(cmd *cobra.Command, args []string) {
	if v, _ := cmd.Flags().GetBool("config"); v {
		ConfigWorkload(cmd, args)
	} else if v, _ := cmd.Flags().GetBool("setup"); v {
		SetupWorkload(cmd, args)
	} else {
		SelectWorkload(cmd, args)
	}
}

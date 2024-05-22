package cmd

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	pflags := rootCmd.PersistentFlags()
	pflags.Bool("log-report-caller", false, "Whenever or not to output the file that outputs the log")
	pflags.String("log-level", "info", "The debug level, levels are: debug, info, warn, error, fatal")
	pflags.String("log-format", "text", "The text format of the logger")
}

func SetupLogger(cmd *cobra.Command, args []string) {
	logOptions := log.Options{
		TimeFormat:   time.DateTime,
		ReportCaller: cmd.Flag("log-report-caller").Value.String() == "true",
	}

	// log-level
	level, err := log.ParseLevel(cmd.Flag("log-level").Value.String())
	if err != nil {
		log.Fatal("invalid log level", "error", err)
	}
	logOptions.Level = level

	// log-format
	switch cmd.Flag("log-format").Value.String() {
	default:
		logOptions.Formatter = log.TextFormatter
	case "json":
		logOptions.Formatter = log.JSONFormatter
	case "logfmt":
		logOptions.Formatter = log.LogfmtFormatter
	}

	// Initialize the default logger.
	logger := log.NewWithOptions(os.Stderr, logOptions)
	log.SetDefault(logger)
}

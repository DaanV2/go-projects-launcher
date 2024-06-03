package cmd

import (
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
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
		ReportTimestamp: false,
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
	logger.SetStyles(CreateStyle())
	log.SetDefault(logger)
}

func CreateStyle() *log.Styles {
	styles := log.DefaultStyles()

	styles.Levels[log.DebugLevel] = styles.Levels[log.DebugLevel].SetString("🔎")
	styles.Levels[log.InfoLevel] = styles.Levels[log.InfoLevel].SetString("🚀")
	styles.Levels[log.WarnLevel] = styles.Levels[log.WarnLevel].SetString("⚠️")
	styles.Levels[log.ErrorLevel] = styles.Levels[log.ErrorLevel].SetString("💥")
	styles.Levels[log.FatalLevel] = styles.Levels[log.FatalLevel].SetString("☠️")

	styles.Keys["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	styles.Keys["error"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	styles.Values["error"] = lipgloss.NewStyle().Bold(true)
	styles.Values["error"] = lipgloss.NewStyle().Bold(true)
	
	return styles
}
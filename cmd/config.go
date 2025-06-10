package cmd

import (
	"path/filepath"

	"github.com/DaanV2/go-projects-launcher/pkg/config"
	"github.com/DaanV2/go-projects-launcher/pkg/ide"
	"github.com/DaanV2/go-projects-launcher/pkg/projects"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func ConfigWorkload(cmd *cobra.Command, args []string) {
	fp := config.ConfigFilepath()
	log.Info("Configuration file path", "path", fp)

	dir := filepath.Dir(fp)

	userIde, _ := cmd.Flags().GetString("ide")
	pro := &projects.Project{
		Name:   "Tool Config",
		Folder: dir,
	}

	err := invokeIDE(ide.GetIDE(userIde), pro, config.GetConfig())
	if err != nil {
		log.Fatal("error while invoking IDE", "error", err)
	}
}

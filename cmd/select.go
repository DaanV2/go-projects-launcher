package cmd

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/DaanV2/go-projects-launcher/pkg/config"
	"github.com/DaanV2/go-projects-launcher/pkg/ide"
	"github.com/DaanV2/go-projects-launcher/pkg/projects"
	"github.com/DaanV2/go-projects-launcher/pkg/regex"
	"github.com/DaanV2/go-projects-launcher/pkg/slicesx"
	"github.com/spf13/cobra"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
)

// SelectWorkload opens a project in the selected IDE
func SelectWorkload(cmd *cobra.Command, args []string) {
	var selectProject *projects.Project
	c := config.GetConfig()
	projs := projects.GetProjects(c.Match, c.ProjectFolders)

	// If filter pattern specified apply
	if len(args) > 0 {
		filter := args[0]
		projs = slicesx.Filter(projs, func(item *projects.Project) bool {
			return regex.IsMatch(c.Match, item.Folder, filter)
		})
	}

	projOptions := make([]huh.Option[projects.Project], 0, len(projs))
	for _, proj := range projs {
		projOptions = append(projOptions, huh.NewOption(proj.Name, *proj))
	}

	// If there are projects left to select, make a form
	if len(projs) > 1 {
		selectProject = projs[0]
		selectForm := huh.NewSelect[projects.Project]().
			Title("Select project").
			Description("The project to open").
			Options(projOptions...).
			Value(selectProject)

		if FatalIfNotAbort(Display("Project launcher", "", selectForm)) {
			return
		}
	} else if len(projs) == 1 {
		selectProject = projs[0]
	}

	if selectProject == nil {
		log.Fatal("no project selected")
		return
	}

	userIde, _ := cmd.Flags().GetString("ide")
	err := invokeIDE(ide.GetIDE(userIde), selectProject, c)
	if err != nil {
		log.Fatal("trouble launching the ide", "error", err)
	}
}

func invokeIDE(ideC ide.IDE, project *projects.Project, userConfig *config.Config) error {
	logger := log.WithPrefix("Launching")

	var c *config.IDEConfig
	if ideC == nil {
		ideC, c = findIDE(project, userConfig)
	}
	if ideC == nil {
		return errors.New("couldn't find a IDE to launch this project for")
	}

	ocom := ideC.OpenCommand(project.Folder)
	if c != nil {
		switch ideC.ID() {
		case ide.CUSTOM:
			cmd := strings.ReplaceAll(c.Custom, "{folder}", project.Folder)
			ocom = exec.Command(cmd)
		case ide.CUSTOM_WLS:
			cmd := strings.ReplaceAll(c.Custom, "{folder}", project.Folder)
			ocom = exec.Command("wsl", cmd)
		}
	}
	logger.Debug("Attempting...",
		"command", ocom,
		"project", project.Name,
		"folder", project.Folder,
	)
	if ocom == nil {
		return errors.New("couldn't find a command to launch this project for")
	}

	ocom.Dir = project.Folder
	logger.Info(project.Name + " 🌕")
	if err := ocom.Start(); err != nil {
		return err
	}

	return nil
}

func findIDE(project *projects.Project, config *config.Config) (ide.IDE, *config.IDEConfig) {
	for _, i := range config.IDE {
		if regex.IsMatch(config.Match, project.Folder, i.PathFilter) {
			return i.IDE.Get(), i
		}
	}

	return config.DefaultIDE.Get(), nil
}

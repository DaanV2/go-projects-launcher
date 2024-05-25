package cmd

import (
	"github.com/DaanV2/go-projects-launcher/config"
	"github.com/DaanV2/go-projects-launcher/ide"
	"github.com/DaanV2/go-projects-launcher/slicesx"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func SetupWorkload(cmd *cobra.Command, args []string) {
	c := config.GetConfig()
	ides := ide.GetIDEs()
	selectedIde := ide.DetermineIDEFromSystem()
	opts := slicesx.Map(ides, func(i ide.IDE) huh.Option[ide.IDE_ID] {
		return huh.NewOption(i.Name(), i.ID())
	})

	ideSelect := huh.NewMultiSelect[ide.IDE_ID]().
		Title("IDE to configure").
		Description("The IDE that are supported, the ones already selected are already found on the system").
		Value(&selectedIde).
		Options(opts...)

	if ErrorIfNotAbort(Display("Configuration", "", ideSelect)) {
		return
	}
	if len(selectedIde) > 0 {
		c.DefaultIDE = selectedIde[0]
		for _, i := range selectedIde {
			c.IDE = append(c.IDE, configIDE(i))
		}
	}

	config.SaveConfig(c)
	log.Infof("save file made: '%s'", config.ConfigFilepath())
}

func configIDE(i ide.IDE_ID) *config.IDEConfig {
	ideConfig := &config.IDEConfig{
		IDE:        i,
		PathFilter: "*",
	}
	title := i.String()

	opts := []huh.Option[string]{
		huh.NewOption("*", ""),
	}

	if data := i.Get(); data != nil {
		title = data.Name()
		opts = append(
			opts,
			slicesx.Map(data.RecommendPatterns(), func(item string) huh.Option[string] {
				return huh.NewOption(item, item)
			})...,
		)
	}

	selectForm := huh.NewSelect[string]().
		Title("Select pattern").
		Description("The regex pattern to use to determine if this IDE should be used on what project").
		Options(opts...).
		Value(&ideConfig.PathFilter)

	ErrorIfNotAbort(Display(title, "Configuration of this IDE in the config", selectForm))
	return ideConfig
}

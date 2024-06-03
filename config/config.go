package config

import (
	"os"
	"path"

	"github.com/DaanV2/go-projects-launcher/ide"
	"github.com/DaanV2/go-projects-launcher/users"
	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

type (
	// Config is the configuration for the projects switcher
	Config struct {
		DefaultIDE     ide.IDE_ID       `yaml:"default_ide"` // The default IDE to use
		ProjectFolders []*ProjectFolder `yaml:"project_folders"` // The folders to look through
		IDE            []*IDEConfig     `yaml:"ide_config"` // The IDE to configure
	}

	// IDEConfig is the configuration for an IDE
	IDEConfig struct {
		IDE        ide.IDE_ID `yaml:"ide"` // The IDE to configure
		PathFilter string     `yaml:"path_filter"` // The regex filter to apply, if blank then always true
		Custom     string     `yaml:"custom"`      // A custom command to run, if blank then use the default command
	}

	// ProjectFolder is the configuration for a folder to look through
	ProjectFolder struct {
		Folder   string   `yaml:"folder"`   // The folder to look through
		Includes []string `yaml:"includes"` // The regex pattern that must match an item
		Excludes []string `yaml:"excludes"` // The regex pattern that when match will excluded that item
	}
)

// ConfigFilepath returns the path to the config file
func ConfigFilepath() string {
	c := users.ConfigDirectory()
	toolFolder := path.Join(c, ".projects")
	logger := log.WithPrefix("Config").With("folder", toolFolder)

	if _, err := os.Stat(toolFolder); err != nil {
		logger.Debug("creating tool folder", "folder", toolFolder)

		if err := os.MkdirAll(toolFolder, os.ModePerm); err != nil {
			logger.Fatal("could not create folder for tool", "error", err, "folder", toolFolder)
		}
	}

	return path.Join(toolFolder, "projects-switch.yaml")
}

// GetConfig returns the configuration for the projects switcher
func GetConfig() *Config {
	filepath := ConfigFilepath()

	c, err := load(filepath)
	logger := log.WithPrefix("Config").With("filepath", filepath)

	if os.IsNotExist(err) {
		logger.Info("No config file found, make sure to go through the --config or --setup steps, or edit the manual file")
		SaveConfig(c)
	} else if err != nil {
		logger.Error("error loading config file", "error", err)
	}

	return c
}

// SaveConfig saves the configuration for the projects switcher
func SaveConfig(c *Config) {
	if err := save(ConfigFilepath(), c); err != nil {
		log.WithPrefix("Config").Error("error while saving", "error", err)
	}
}

// DefaultConfig returns the default configuration for the projects switcher
func DefaultConfig() *Config {
	return &Config{
		ProjectFolders: []*ProjectFolder{
			{
				Folder:   path.Join(users.HomeDirectory(), "repos"),
				Includes: []string{},
				Excludes: []string{},
			},
		},
		IDE:        []*IDEConfig{},
		DefaultIDE: ide.VSCODE,
	}
}

// load the configuration from a file
func load(filepath string) (*Config, error) {
	log.WithPrefix("Config").Debug("loading...", "file", filepath)

	c := DefaultConfig()
	data, err := os.ReadFile(filepath)
	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(data, c)
	return c, err
}

// save the configuration to a file
func save(filepath string, c *Config) error {
	log.WithPrefix("Config").Debug("saving...", "file", filepath)

	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, data, 0644)
}

package config

import (
	"os"
	"path"

	"github.com/DaanV2/projects-tool/ide"
	"github.com/DaanV2/projects-tool/users"
	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		DefaultIDE     ide.IDE_ID       `yaml:"default_ide"`
		ProjectFolders []*ProjectFolder `yaml:"project_folders"`
		IDE            []*IDEConfig     `yaml:"ide_config"`
	}

	IDEConfig struct {
		IDE        ide.IDE_ID `yaml:"ide"`
		PathFilter string     `yaml:"path_filter"` // The regex filter to apply, if blank then always true
	}

	ProjectFolder struct {
		Folder   string   `yaml:"folder"`   // The folder to look through
		Includes []string `yaml:"includes"` // The regex pattern that must match an item
		Excludes []string `yaml:"excludes"` // The regex pattern that when match will excluded that item
	}
)

func ConfigFilepath() string {
	c := users.ConfigDirectory()

	toolFolder := path.Join(c, ".projects")
	if _, err := os.Stat(toolFolder); err != nil {
		log.Debug("creating tool folder", "folder", toolFolder)

		if err := os.MkdirAll(toolFolder, os.ModePerm); err != nil {
			log.Fatal("could not create folder for tool", "error", err, "folder", toolFolder)
		}
	}

	return path.Join(toolFolder, "projects-switch.yaml")
}

func GetConfig() *Config {
	filepath := ConfigFilepath()

	c, err := load(filepath)
	if os.IsNotExist(err) {
		log.Infof("No config file found, make sure to go through the --config or --setup steps, or edit the manual file: %s", filepath)
		SaveConfig(c)
	} else if err != nil {
		log.Errorf("error loading config file: %v", err)
	}

	return c
}

func SaveConfig(c *Config) {
	if err := save(ConfigFilepath(), c); err != nil {
		log.Errorf("error saving config file: %v", err)
	}
}

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

func load(filepath string) (*Config, error) {
	log.Debugf("loading config file: %s", filepath)

	c := DefaultConfig()
	data, err := os.ReadFile(filepath)
	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(data, c)
	return c, err
}

func save(file string, c *Config) error {
	log.Debugf("saving config file: %s", file)

	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0644)
}

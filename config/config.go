package config

import (
	"os"
	"path"
	"path/filepath"

	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ProjectFolders []string `yaml:"project_folders"`
	IDE []IDEConfig `yaml:"ide_config"`
}

type IDEConfig struct {
	Command    string `yaml:"path_filter"`
	PathFilter string `yaml:"path_filter"` // The regex filter to apply, if blank then always true
}

var config_filepath = path.Join("~", ".projects", "projects-switch.yaml")

func GetConfig() *Config {
	c, err := load(config_filepath)
	if err != nil {
		log.Errorf("error loading config file: %v", err)
		SaveConfig(c)
	}

	return c
}

func SaveConfig(c *Config) {
	if err := save(config_filepath, c); err != nil {
		log.Errorf("error saving config file: %v", err)
	}
}

func DefaultConfig() *Config {
	return &Config{
		ProjectFolders: []string{
			path.Join("~", "repos"),
		},
		IDE: []IDEConfig{},
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

	directory := filepath.Dir(file)
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if err := os.Mkdir(directory, os.ModePerm); err != nil {
			log.Errorf("could not create folder: %v", err)
			return err
		}
	}

	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0644)
}
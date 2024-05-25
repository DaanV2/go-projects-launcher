package users

import (
	"os"
	"path"

	"github.com/charmbracelet/log"
)

// CacheDirectory returns the cache directory of the user
func CacheDirectory() string {
	c, err := os.UserCacheDir()
	if err != nil {
		log.Error("cannot determine the users config directory", "error", err)
		c = path.Join(".", ".cache")
	}

	return c
}

// HomeDirectory returns the user home directory
func HomeDirectory() string {
	c, err := os.UserHomeDir()
	if err != nil {
		log.Error("cannot determine the users config directory", "error", err)
		c = path.Join(".", ".home")
	}

	return c
}

// ConfigDirectory returns the config directory of the user
func ConfigDirectory() string {
	c, err := os.UserConfigDir()
	if err != nil {
		log.Error("cannot determine the users config directory", "error", err)
		c = path.Join(".", ".config")
	}

	return c
}

package users

import (
	"os"
	"path"

	"github.com/charmbracelet/log"
)

func CacheDirectory() string {
	c, err := os.UserCacheDir()
	if err != nil {
		log.Error("cannot determine the users config directory", "error", err)
	} else {
		c = path.Join(".", ".cache")
	}

	return c
}

func HomeDirectory() string {
	c, err := os.UserHomeDir()
	if err != nil {
		log.Error("cannot determine the users config directory", "error", err)
	} else {
		c = path.Join(".", ".home")
	}

	return c
}

func ConfigDirectory() string {
	c, err := os.UserConfigDir()
	if err != nil {
		log.Error("cannot determine the users config directory", "error", err)
	} else {
		c = path.Join(".", ".config")
	}

	return c
}
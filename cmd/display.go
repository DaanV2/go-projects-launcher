package cmd

import (
	"errors"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
)

func Display(title string, description string, fields ...huh.Field) error {
	group := huh.NewGroup(fields...)
	if title != "" {
		group.Title(title)
	}
	if description != "" {
		group.Description(description)
	}
	form := huh.NewForm(group)
	return form.Run()
}

func FatalIfNotAbort(err error) bool {
	if errors.Is(err, huh.ErrUserAborted) {
		return true
	}
	if err == nil {
		return false
	}

	log.Fatal("error while executing forms", "error", err)
	return false
}

// ErrorIfNotAbort return true is aborted
func ErrorIfNotAbort(err error) bool {
	if errors.Is(err, huh.ErrUserAborted) {
		return true
	}
	if err == nil {
		return false
	}

	log.Error("error while executing forms", "error", err)
	return false
}
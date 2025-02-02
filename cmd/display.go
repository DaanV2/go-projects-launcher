package cmd

import (
	"errors"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
)

// Display wraps the given fields into a group and form and shows it, if title and/or description are not "", then they will be configured and add to the form
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

// FatalIfNotAbort Checks if the error is an "user aborted" error and return true if it is, Else will fataly log the error if not nil
func FatalIfNotAbort(err error) bool {
	if errors.Is(err, huh.ErrUserAborted) {
		log.Info("Bye 👋")
		return true
	}
	if err == nil {
		return false
	}

	log.Fatal("error while executing forms", "error", err)
	return false
}

// ErrorIfNotAbort Checks if the error is an "user aborted" error and return true if it is, Else will error log the error if not nil
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

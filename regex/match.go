package regex

import (
	go_regexp "regexp"

	"github.com/charmbracelet/log"
)

// IsMatch checks the given item against all the provided arguments, expected that the patterns is a regex string
func IsMatch(item string, patterns ...string) bool {
	for _, pattern := range patterns {
		matched, err := go_regexp.MatchString(pattern, item)
		if err != nil {
			log.Error("error with regex pattern", "pattern", pattern, "error", err)
		}
		if matched {
			return true
		}
	}

	return false
}

// Determine check if an item matches with the include and not with the excludes, if either one is 0
func Determine(item string, include []string, exclude []string) bool {
	result := true

	if len(include) > 0 {
		result = IsMatch(item, include...)
	}
	if result && len(exclude) > 0 {
		result = !IsMatch(item, exclude...)
	}

	return result
}

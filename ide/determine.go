package ide

import (
	"os"
	"slices"
	"strings"

	"github.com/DaanV2/projects-tool/slicesx"
)

// DetermineIDEFromSystem determines from the system what types of IDE are installed
func DetermineIDEFromSystem() []IDE_ID {
	found := make([]IDE_ID, 0)

	if term := os.Getenv("TERM_PROGRAM"); len(term) > 0 {
		if item := GetIDE(term); item != nil {
			found = append(found, item.ID())
		}
	}

	environs := os.Environ()
	// Vscode?
	if slices.ContainsFunc(environs, func(item string) bool {
		return strings.HasPrefix(item, "VSCODE")
	}) {
		found = append(found, VSCODE, VSCODE_WLS)
	}

	// Only include IDE that are registered. (wls for example is not register if not windows)
	found = slicesx.Filter(found, func(item IDE_ID) bool {
		return item.Get() != nil
	})
	return slicesx.Unique(found)
}

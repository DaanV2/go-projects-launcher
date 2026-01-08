package ide

import (
	"os/exec"
	"runtime"
)

const (
	GOLAND   IDE_ID = "goland"
	INTELLIJ IDE_ID = "intellij"
)

func init() {
	registerIDE(
		ideConfig{
			func(folder string) *exec.Cmd {
				return exec.Command("idea", ".")
			},
			"Goland - Jetbrains",
			GOLAND,
			[]string{},
		},
		ideConfig{
			func(folder string) *exec.Cmd {
				return exec.Command("idea", ".")
			},
			"Intellij - Jetbrains",
			INTELLIJ,
			[]string{},
		},
	)
}
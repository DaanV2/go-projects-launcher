package ide

import (
	"os/exec"
	"runtime"
)

const (
	VSCODE     IDE_ID = "vscode"
	VSCODE_WLS IDE_ID = "wsl-vscode"
)

func init() {
	registerIDE(
		ideConfig{
			func(folder string) *exec.Cmd {
				return exec.Command("code", ".")
			},
			"Visual Studio Code",
			VSCODE,
			[]string{},
		},
	)

	if runtime.GOOS == "windows" {
		registerIDE(
			ideConfig{
				func(folder string) *exec.Cmd {
					return exec.Command("wsl", "code", ".")
				},
				"Visual Studio Code (WSL)",
				VSCODE_WLS,
				[]string{"wsl.localhost"},
			},
		)
	}
}

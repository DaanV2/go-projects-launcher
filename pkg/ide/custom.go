package ide

import (
	"os/exec"
	"runtime"
)

const (
	CUSTOM     IDE_ID = "custom"
	CUSTOM_WLS IDE_ID = "wsl-custom"
)

func init() {
	registerIDE(
		ideConfig{
			func(folder string) *exec.Cmd {
				return nil
			},
			"Custom",
			CUSTOM,
			[]string{},
		},
	)

	if runtime.GOOS == "windows" {
		registerIDE(
			ideConfig{
				func(folder string) *exec.Cmd {
					return exec.Command("wsl")
				},
				"Custom (WSL)",
				CUSTOM_WLS,
				[]string{"wsl.localhost"},
			},
		)
	}
}

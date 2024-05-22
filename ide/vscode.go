package ide

const (
	VSCODE     IDE_ID = "vscode"
	VSCODE_WLS IDE_ID = "wsl-vscode"
)

func init() {
	registerIDE(
		ideConfig{
			"code",
			"Visual Studio Code",
			VSCODE,
			[]string{"*"},
		},
		ideConfig{
			"wsl code .",
			"Visual Studio Code (WSL)",
			VSCODE_WLS,
			[]string{"wsl.localhost"},
		},
	)
}

package ide

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

package ide

const (
	INTELLIJ      IDE_ID = "intellij"
)

func init() {
	registerIDE(
		ideConfig{
			func(folder string) *exec.Cmd {
				return exec.Command("idea", ".")
			},
			"intellij",
			INTELLIJ,
			[]string{},
		},
	)
}
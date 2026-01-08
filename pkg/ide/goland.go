package ide

const (
	GOLAND IDE_ID = "goland"
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
	)
}
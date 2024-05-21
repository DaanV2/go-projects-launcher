package ide

import (
	"fmt"
	"maps"
)

type (
	IDE_ID string

	IDE interface {
		// OpenCommand returns the command to open the IDE
		OpenCommand() string
		// Name returns the name of the IDE
		Name() string
		// ID returns the ID of the IDE
		ID() IDE_ID
		// RecommendPatterns returns a list of repo patterns that the IDE is recommended for
		RecommendPatterns() []string
	}

	ideConfig struct {
		_OpenCommand string
		_Name        string
		_ID          IDE_ID
		_Recommend   []string
	}
)

func (i ideConfig) OpenCommand() string { return i._OpenCommand }
func (i ideConfig) Name() string        { return i._Name }
func (i ideConfig) ID() IDE_ID          { return i._ID }
func (i ideConfig) RecommendPatterns() []string { return i._Recommend }

var ides = map[IDE_ID]IDE{}

func registerIDE(ide ...IDE) {
	for _, i := range ide {
		if _, ok := ides[i.ID()]; ok {
			panic(fmt.Errorf("IDE %s already registered", i.Name()))
		} else {
			ides[i.ID()] = i
		}
	}
}

// GetIDE returns the IDE with the given ID, or false if it does not exist
func GetIDE(id IDE_ID) (IDE, bool) {
	i, ok := ides[id]
	return i, ok
}

// GetIDEs returns a copy of the map of IDEs
func GetIDEs() map[IDE_ID]IDE {
	return maps.Clone(ides)
}

const (
	VSCODE     IDE_ID = "vscode"
	VSCODE_WLS IDE_ID = "wsl-vscode"
)

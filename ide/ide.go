package ide

import (
	"log"
)

type (
	IDE_ID string

	IDE interface {
		// OpenCommand returns the command to open the IDE
		OpenCommand(folder string) string
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

func (i IDE_ID) String() string                      { return string(i) }
func (i IDE_ID) Get() IDE                            { return GetIDE(i.String()) }
func (i ideConfig) OpenCommand(folder string) string { return i._OpenCommand }
func (i ideConfig) Name() string                     { return i._Name }
func (i ideConfig) ID() IDE_ID                       { return i._ID }
func (i ideConfig) RecommendPatterns() []string      { return i._Recommend }

var ides = map[IDE_ID]IDE{}

func registerIDE(ide ...IDE) {
	for _, i := range ide {
		if ide == nil {
			log.Fatal("ide specification is nil")
		}

		if _, ok := ides[i.ID()]; ok {
			log.Fatalf("IDE %s already registered", i.Name())
		} else {
			ides[i.ID()] = i
		}
	}
}

// GetIDE returns the IDE with the given ID, or false if it does not exist
func GetIDE(id string) IDE {
	idx := IDE_ID(id)
	if i, ok := ides[idx]; ok {
		return i
	}
	return nil
}

// GetIDEs returns a copy of the map of IDEs
func GetIDEs() []IDE {
	result := make([]IDE, 0)

	for _, i := range ides {
		result = append(result, i)
	}

	return result
}

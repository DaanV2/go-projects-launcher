package projects

import (
	"os"
	"path"

	"github.com/DaanV2/go-projects-launcher/pkg/config"
	"github.com/DaanV2/go-projects-launcher/pkg/regex"
	"github.com/charmbracelet/log"
)

var cache = make(map[string][]*Project)

type Project struct {
	Folder string
	Name   string
}

func GetProjects(folders []*config.ProjectFolder) []*Project {
	result := make([]*Project, 0, 100)

	for _, folder := range folders {
		if items, ok := cache[folder.Folder]; ok {
			result = append(result, items...)
			continue
		}
		if items, err := discoverProjects(folder); err != nil {
			delete(cache, folder.Folder)
		} else {
			cache[folder.Folder] = items
			result = append(result, items...)
		}
	}

	return result
}

// discoverProjects finds all the projects in the given folder
func discoverProjects(folder *config.ProjectFolder) ([]*Project, error) {
	log.Debug("discovering project folder", "folder", folder.Folder)
	result := make([]*Project, 0, 100)

	items, err := os.ReadDir(folder.Folder)
	if err != nil {
		log.Error("couldn't check folder for project", "folder", folder.Folder)
		return result, err
	}

	for _, item := range items {
		if !item.IsDir() {
			continue
		}
		fullpath := path.Join(folder.Folder, item.Name())

		if regex.Determine(fullpath, folder.Includes, folder.Excludes) {
			result = append(result, &Project{
				Folder: fullpath,
				Name:   item.Name(),
			})
		}
	}

	return result, nil
}

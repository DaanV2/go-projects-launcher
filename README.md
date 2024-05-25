# Projects tool

A tool for launching projects, allowing you to search all your different projects and launch them. You configure the tool with where your projects are located and it will search for them and allow you to launch them. The program that will be used to launch the project is configurable. With filters, you can dynamically filter the projects that you want to launch.

```bash
go install github.com/DaanV2/go-projects-launcher@v1.0.0
```

## Examples

```bash
go-projects-launcher some-project-folder-pattern
```

## Usage

```bash
# Setup
go-projects-launcher --setup

go-projects-launcher [flags] [project-folder-pattern]

# If only 1 project is found, it will be launched
# else a list of projects will be shown
go-projects-launcher projectA
go-projects-launcher proj

# Show all
go-projects-launcher
```

## Configure

```yaml
default_ide: vscode
project_folders:
  - folder: /workspaces/<user>/.home/repos
    includes: [] # regex patterns that must match
    excludes: [] # regex patterns that must not match
ide_config:
  - ide: vscode-wsl
    path_filter: wsl.localhost # regex pattern that must match to use this ide
  - ide: vscode
    path_filter: ""
```
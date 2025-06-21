package main

import (
	"context"

	"github.com/DaanV2/go-projects-launcher/cmd"
	"github.com/charmbracelet/fang"
	"github.com/charmbracelet/log"
)

func main() {
	err := fang.Execute(context.Background(), cmd.RootCMD())
	if err != nil {
		log.Fatal("error executing command: %v", err)
	}
}

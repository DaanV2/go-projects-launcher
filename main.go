package main

import (
	"github.com/DaanV2/projects-tool/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	cmd.Execute()
}

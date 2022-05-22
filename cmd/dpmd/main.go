package main

import (
	"flag"

	"github.com/godpm/godpm/config"
	"github.com/godpm/godpm/pkg/daemon"
)

var (
	conf = flag.String("conf", "godpmd.config.yaml", "dpmd config file")
)

func setupConfig() {
	config.SetupConfiguration(*conf)
}

func main() {
	setupConfig()
	daemon.Start(config.AppConfig.LogFile, config.AppConfig.PidFile)
}

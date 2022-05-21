package main

import (
	"flag"
	"time"

	"github.com/godpm/godpm/config"
	"github.com/godpm/godpm/pkg/daemon"
	"github.com/godpm/godpm/pkg/log"
)

var (
	conf = flag.String("conf", "godpmd.config.yaml", "dpmd config file")
)

func setupConfig() {
	config.SetupConfiguration(*conf)
}

func main() {
	setupConfig()
	daemon.Start(config.AppConfig.LogFile, config.AppConfig.PidFile, func() {
		for i := 20; i > 0; i-- {
			log.Error().Println("idx", 1)
			time.Sleep(1 * time.Second)
		}
	})
}

package main

import (
	"flag"

	"github.com/godpm/godpm/config"
	"github.com/godpm/godpm/pkg/daemon"
	//	"github.com/godpm/godpm/pkg/process"
)

var (
	conf = flag.String("conf", "godpmd.config.yaml", "dpmd config file")
)

func setupConfig() {
	config.SetupConfiguration(*conf)
}

func main() {
	setupConfig()
	daemonStart()
}

func daemonStart() {
	daemon.Start(config.AppConfig.LogFile, config.AppConfig.PidFile)
}

/*
func runProcess() {
	process.InitAndStart()
	select {}
}
*/

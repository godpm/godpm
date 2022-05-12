package main

import (
	"flag"

	"github.com/godpm/godpm/cmd/dpmd/config"
)

var (
	conf = flag.String("conf", "dpmd.config.yaml", "dpmd config file")
)

func setupConfig() {
	config.SetupConfiguration(*conf)
}

func main() {
	setupConfig()
}

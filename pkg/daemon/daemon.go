package daemon

import (
	"github.com/godpm/godpm/pkg/log"
	"github.com/sevlyar/go-daemon"
)

// Start start daemon and process from config dir
func Start(logFile, pidFile string, f func()) {
	context := daemon.Context{PidFileName: pidFile, LogFileName: logFile}

	child, err := context.Reborn()
	if err != nil {
		log.Fatal().Fatal(err.Error())
	}

	if child != nil {
		return
	}

	defer func() {
		if err := context.Release(); err != nil {
			log.Error().Println("daemon context release failed ", err)
		}
	}()

	f()
}

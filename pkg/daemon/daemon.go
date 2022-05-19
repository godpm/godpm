package daemon

import (
	"log"

	"github.com/sevlyar/go-daemon"
)

// Start start daemon and process from config dir
func Start(logFile, pidFile string, f func()) {
	context := daemon.Context{PidFileName: pidFile, LogFileName: logFile}

	child, err := context.Reborn()
	if err != nil {
		log.Fatal(err.Error())
	}

	if child != nil {
		return
	}

	defer context.Release()

	f()
}

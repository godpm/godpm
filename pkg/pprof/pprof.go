package pprof

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/godpm/godpm/config"
	"github.com/godpm/godpm/pkg/log"
)

// Run run pprof
func Run() {
	if config.AppConfig.PprofPort > 0 {
		server := http.Server{
			Addr: fmt.Sprintf(":%d", config.AppConfig.PprofPort),
		}

		if err := server.ListenAndServe(); err != nil {
			log.Fatal().Println("Start pprof failed ", err)
		}
	}
}

package http

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/godpm/godpm/config"
	"github.com/godpm/godpm/pkg/log"
)

// RunServer run http server
func RunServer() {
	log.Fatal().Printf("%v", fasthttp.ListenAndServe(fmt.Sprintf(":%d", config.AppConfig.Port), configRoute().Handler))
}

func configRoute() *router.Router {
	r := router.New()
	r.GET("/", func(ctx *fasthttp.RequestCtx) {})
	v1 := r.Group("/v1")
	v1.POST("/deploy", func(ctx *fasthttp.RequestCtx) {
	})

	v1.PUT("/stop/{name}", Stop)
	v1.GET("/status", Status)
	v1.PUT("/restart/{name}", Restart)
	v1.PUT("/start/{name}", Start)

	return r
}

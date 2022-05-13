package http

import (
	"github.com/fasthttp/router"
	"github.com/phuslu/log"
	"github.com/valyala/fasthttp"
)

// RunServer run http server
func RunServer() {
	log.Fatal().Msgf("%v", fasthttp.ListenAndServe(":8080", configRoute().Handler))
}

func configRoute() *router.Router {
	r := router.New()
	r.GET("/", func(ctx *fasthttp.RequestCtx) {})
	return r
}

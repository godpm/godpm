package http

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// RunServer run http server
func RunServer() {
	fasthttp.ListenAndServe(":8080", configRoute().Handler)
}

func configRoute() *router.Router {
	r := router.New()
	r.GET("/", func(ctx *fasthttp.RequestCtx) {})
	return r
}

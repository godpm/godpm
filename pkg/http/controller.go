package http

import (
	"github.com/valyala/fasthttp"
)

// RunCMDArgs ..
type RunCMDArgs struct {
	Name      string `json:"name"`
	Env       string `json:"env"`
	CMD       string `json:"cmd"`
	Restart   string `json:"restart"` //
	Directory string `json:"directory"`
	PreCMD    string `json:"pre_cmd"` //
	User      string `json:"user"`    //
}

// RunCMD
func RunCMD(ctx *fasthttp.RequestCtx) {
	return
}

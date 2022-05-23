package http

import (
	"encoding/json"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/godpm/godpm/pkg/process"
)

// Stop stop process
func Stop(ctx *fasthttp.RequestCtx) {
	name := ctx.UserValue("name").(string)
	process.Stop(name)
}

// ProcStatus process status
type ProcStatus struct {
	Name   string    `json:"name"`
	State  string    `json:"state"`
	Uptime time.Time `json:"uptime"`
	Pid    int       `json:"pid"`
}

// Status status
func Status(ctx *fasthttp.RequestCtx) {
	var (
		names = ctx.FormValue("names")
		procs = []*process.Process{}
	)

	if len(names) == 0 {
		procs = process.List()
	} else {
		proc, ok := process.Find(string(names))
		if !ok {
			ctx.Write([]byte(`{"message": "process not found"}`))
			return
		}
		procs = append(procs, proc)
	}

	resp := make([]ProcStatus, 0, len(procs))
	for _, p := range procs {
		resp = append(resp, ProcStatus{
			Pid:    p.Pid(),
			Name:   p.Name(),
			State:  p.State(),
			Uptime: p.Uptime(),
		})
	}

	b, _ := json.Marshal(&resp)
	ctx.Write(b)
}

// Restart ...
func Restart(ctx *fasthttp.RequestCtx) {
	name := ctx.UserValue("name").(string)
	process.Restart(name)
}
func Start(ctx *fasthttp.RequestCtx) {
	name := ctx.UserValue("name").(string)
	process.Start(name)
}

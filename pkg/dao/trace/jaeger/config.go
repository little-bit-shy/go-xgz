package jaeger

import (
	"github.com/go-kratos/kratos/pkg/net/trace"
)

type Cfg struct {
	Config *Config
	AppID  string
}

// Init Init
func Init(c *Cfg) {
	trace.SetGlobalTracer(trace.NewTracer(c.AppID, newReport(c.Config), true))
}

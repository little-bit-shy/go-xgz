package job

import (
	"context"
	"github.com/go-kratos/kratos/pkg/net/trace"
	"strconv"
)

type ctxKey string

var _ctxkey ctxKey = "kratos/pkg/net/trace.trace"

const _defaultComponentName = "cli"

type Job struct {
	name string
	Do   func(ctx context.Context) (err error)
}

func New(name string, do func(ctx context.Context) (err error)) *Job {
	return &Job{
		name: name,
		Do:   do,
	}
}

func (j Job) Run() {
	var opts []trace.Option
	if ok, _ := strconv.ParseBool(trace.KratosTraceDebug); ok {
		opts = append(opts, trace.EnableDebug())
	}
	t := trace.New(_defaultComponentName, opts...)
	t.SetTitle(j.name)
	t.SetTag(trace.String(trace.TagComponent, _defaultComponentName))
	t.SetTag(trace.String(trace.TagSpanKind, "server"))
	ctx := context.Background()
	// export trace id.
	ctx = trace.NewContext(ctx, t)
	err := j.Do(ctx)
	t.Finish(&err)
}

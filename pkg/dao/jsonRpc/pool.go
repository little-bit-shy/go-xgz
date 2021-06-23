package jsonRpc

import (
	"context"
	"github.com/go-kratos/kratos/pkg/container/pool"
	"github.com/go-kratos/kratos/pkg/net/trace"
	"io"
)

type JsonGrpcPool struct {
	*pool.Slice
}

// NewPool new json rpc client pool
func NewPool(config *Config) (jp *JsonGrpcPool, cf func(), err error) {
	ps := pool.NewSlice(config.Pool)

	// new pool
	ps.New = func(ctx context.Context) (io.Closer, error) {
		var client *JsonRpc
		if t, ok := trace.FromContext(ctx); ok {
			var internalTags []trace.Tag
			internalTags = append(internalTags, trace.TagString(trace.TagComponent, "rpc/jrpc"))
			internalTags = append(internalTags, trace.TagString(trace.TagPeerService, "rpc"))
			internalTags = append(internalTags, trace.TagString(trace.TagSpanKind, "client"))

			t = t.Fork(_family, "Rpc:connect")
			t.SetTag(trace.String(trace.TagAddress, config.Address))
			defer t.Finish(&err)
		}
		client, err = New(config)
		return client, err
	}
	jp = &JsonGrpcPool{ps}
	cf = func() {
		jp.Close()
	}
	return
}

// G get a pool with json rpc
func (j *JsonGrpcPool) G(ctx context.Context) (c *JsonRpc, cf func(), err error) {
	var client io.Closer
	cf = func() {
	}
	if client, err = j.Get(ctx); err != nil {
		return
	}
	if client != nil {
		cf = func() {
			j.Put(ctx, client, false)
		}
	}
	c = client.(*JsonRpc)
	return
}

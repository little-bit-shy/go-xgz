package hbase

import (
	"context"
	"github.com/go-kratos/kratos/pkg/container/pool"
	"github.com/go-kratos/kratos/pkg/net/trace"
	"github.com/little-bit-shy/go-xgz/pkg/database/hbase"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

// Log
type Log struct {
	Enable bool
}

// Config
type Config struct {
	pool *pool.Config
}

// Pool
type Pool struct {
	*pool.Slice
	c Config
}

// NewHbase new hbase pools
func NewHbase(log Log, cfg hbase.Config, poolConfig pool.Config, zkCfg hbase.ZKConfig) (p *Pool, cf func(), err error) {
	if log.Enable == true {
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)
	}
	cfg.Zookeeper = &zkCfg
	conf := Config{
		pool: &poolConfig,
	}
	cPool := Pool{
		c: conf,
	}

	ps := pool.NewSlice(cPool.c.pool)

	// new pool
	ps.New = func(ctx context.Context) (client io.Closer, err error) {
		if t, ok := trace.FromContext(ctx); ok {
			var internalTags []trace.Tag
			internalTags = append(internalTags, trace.TagString(trace.TagComponent, "database/hbase"))
			internalTags = append(internalTags, trace.TagString(trace.TagPeerService, "hbase"))
			internalTags = append(internalTags, trace.TagString(trace.TagSpanKind, "client"))

			t = t.Fork("hbase_client", "Hbase:connect")
			t.SetTag(trace.String(trace.TagAddress, strings.Join(cfg.Zookeeper.Addrs, ",")))
			defer t.Finish(&err)
		}
		client = hbase.NewClient(&cfg)
		return
	}
	p = &Pool{ps, conf}
	cf = func() {
		p.Close()
	}
	return
}

// G get pool
func (h *Pool) G(ctx context.Context) (c *hbase.Client, cf func(), err error) {
	var client io.Closer
	cf = func() {
	}
	if client, err = h.Get(ctx); err != nil {
		return
	}
	if client != nil {
		cf = func() {
			h.Put(ctx, client, false)
		}
	}
	c = client.(*hbase.Client)
	return
}

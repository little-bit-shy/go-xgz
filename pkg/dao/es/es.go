package es

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/pkg/container/pool"
	"github.com/go-kratos/kratos/pkg/net/trace"
	"github.com/little-bit-shy/go-xgz/pkg/elastic"
	"github.com/little-bit-shy/go-xgz/pkg/helper"
	"io"
)

const (
	_family = "jrpc_client"
)

// Cfg es config
type Cfg struct {
	URL       []string
	Sniff     bool
	UserName  string
	PassWorld string
}

// Config pool config
type Config struct {
	pool *pool.Config
	cfg  *Cfg
}

// Pool pool
type Pool struct {
	*pool.Slice
	c Config
}

// Es es client
type Es struct {
	Client *elastic.Client
	io.Closer
}

// NewEs new es pools
func NewEs(PoolConfig pool.Config, cfg Cfg) (p *Pool, cf func(), err error) {
	conf := Config{
		pool: &PoolConfig,
		cfg:  &cfg,
	}
	cPool := Pool{
		c: conf,
	}

	ps := pool.NewSlice(cPool.c.pool)

	// new pool
	ps.New = func(ctx context.Context) (io.Closer, error) {
		es := new(Es)
		if t, ok := trace.FromContext(ctx); ok {
			var internalTags []trace.Tag
			internalTags = append(internalTags, trace.TagString(trace.TagComponent, "databases/elastic"))
			internalTags = append(internalTags, trace.TagString(trace.TagPeerService, "elastic"))
			internalTags = append(internalTags, trace.TagString(trace.TagSpanKind, "client"))

			t = t.Fork(_family, "Elastic:connect")
			t.SetTag(trace.String(trace.TagAddress, fmt.Sprintf("%+v", cfg.URL)))
			defer t.Finish(&err)
		}
		es.NewClient(&cfg)
		return es, nil
	}
	p = &Pool{ps, conf}
	cf = func() {
		p.Close()
	}
	return
}

// NewClient create new es client
func (e *Es) NewClient(cfg *Cfg) {
	client, err := elastic.NewClient(
		elastic.SetURL(cfg.URL...),
		elastic.SetSniff(cfg.Sniff),
		elastic.SetBasicAuth(cfg.UserName, cfg.PassWorld),
	)
	helper.Panic(err)
	e.Client = client
	return
}

// Close close es client
func (e *Es) Close() error {
	return nil
}

// G get pool
func (p *Pool) G(ctx context.Context) (c *elastic.Client, cf func(), err error) {
	var client io.Closer
	cf = func() {
	}

	if client, err = p.Get(ctx); err != nil {
		return
	}
	if client != nil {
		cf = func() {
			p.Put(ctx, client, false)
		}
	} else {
		err = errors.New("es connect error")
		return
	}
	c = client.(*Es).Client

	return
}

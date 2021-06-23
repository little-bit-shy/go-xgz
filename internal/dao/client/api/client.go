package api

import (
	"context"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	pool2 "github.com/go-kratos/kratos/pkg/container/pool"
	service "github.com/little-bit-shy/go-xgz/pkg/api"
	config2 "github.com/little-bit-shy/go-xgz/pkg/config"
	"github.com/little-bit-shy/go-xgz/pkg/helper"
	"github.com/little-bit-shy/go-xgz/pkg/rpc/warden"
	"io"
)

// Pool
type config struct {
	pool *pool2.Config
	conf *warden.ClientConfig
	cc   service.Config
}

type Pool struct {
	*pool2.Slice
	c config
}

// getConfig get config
func getConfig() (conf config, cPool Pool, err error) {
	var (
		cConf        service.Config
		PoolConfig   pool2.Config
		WardenConfig warden.ClientConfig
		ct           paladin.TOML
	)
	err = paladin.Get("client.toml").Unmarshal(&ct)
	helper.Panic(err)
	err = config2.Env(&ct, &cConf, "Api")
	helper.Panic(err)
	err = config2.Env(&ct, &PoolConfig, "Pool")
	helper.Panic(err)
	err = config2.Env(&ct, &WardenConfig, "Warden")
	helper.Panic(err)
	conf = config{
		pool: &PoolConfig,
		conf: &WardenConfig,
		cc:   cConf,
	}
	cPool = Pool{
		c: conf,
	}
	return conf, cPool, err
}

// NewPool new Pool
func NewPool() (p *Pool, cf func(), err error) {
	var conf config
	var cPool Pool
	if conf, cPool, err = getConfig(); err != nil {
		return
	}
	ps := pool2.NewSlice(cPool.c.pool)

	// new pool
	ps.New = func(ctx context.Context) (io.Closer, error) {
		client := service.New(conf.cc.AppId, cPool.c.conf)
		return client, nil
	}
	p = &Pool{ps, conf}
	cf = func() {
		p.Close()
	}
	return
}

// G get pool
func (p *Pool) G(ctx context.Context) (c *service.C, cf func(), err error) {
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
	}
	c = client.(*service.C)
	if c.Err != nil {
		err = c.Err
		return
	}
	return
}

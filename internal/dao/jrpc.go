package dao

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	pool2 "github.com/go-kratos/kratos/pkg/container/pool"
	"github.com/little-bit-shy/go-xgz/pkg/config"
	"github.com/little-bit-shy/go-xgz/pkg/dao/jsonRpc"
	"github.com/little-bit-shy/go-xgz/pkg/helper"
)

type Jp struct {
	Jp  *jsonRpc.JsonGrpcPool
	cf  func()
	err error
}

type Jrpc struct {
	Dict *Jp
}

// NewJrpcPool new all jrpc pool
func NewJrpcPool() (jrpc *Jrpc, cf func(), err error) {
	var jp *jsonRpc.JsonGrpcPool

	jp, cf, err = JrpcPool("Dict", "DictPool")
	helper.Panic(err)
	jrpc = &Jrpc{Dict: &Jp{
		Jp:  jp,
		cf:  cf,
		err: err,
	}}
	cf = func() {
		jrpc.Dict.Jp.Close()
		return
	}
	return
}

// JrpcPool jrpc pool
func JrpcPool(service string, pool string) (jp *jsonRpc.JsonGrpcPool, cf func(), err error) {
	var (
		cfg jsonRpc.Config
		p   pool2.Config
		ct  paladin.Map
	)
	err = paladin.Get("jrpc.toml").Unmarshal(&ct)
	helper.Panic(err)
	err = config.Env(&ct, &cfg, service)
	helper.Panic(err)
	err = config.Env(&ct, &p, pool)
	helper.Panic(err)
	cfg.Pool = &p
	jp, cf, err = jsonRpc.NewPool(&cfg)
	return
}

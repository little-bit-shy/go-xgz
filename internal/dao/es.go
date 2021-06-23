package dao

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/container/pool"
	"github.com/little-bit-shy/go-xgz/pkg/config"
	"github.com/little-bit-shy/go-xgz/pkg/dao/es"
	"github.com/little-bit-shy/go-xgz/pkg/helper"
)

// NewEs new es pools
func NewEs() (p *es.Pool, cf func(), err error) {
	var (
		cfg        es.Cfg
		PoolConfig pool.Config
		ct         paladin.TOML
	)
	err = paladin.Get("es.toml").Unmarshal(&ct)
	helper.Panic(err)
	err = config.Env(&ct, &cfg, "Client")
	helper.Panic(err)
	err = config.Env(&ct, &PoolConfig, "Pool")
	helper.Panic(err)

	return es.NewEs(PoolConfig, cfg)
}

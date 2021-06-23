package dao

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/container/pool"
	"github.com/little-bit-shy/go-xgz/pkg/config"
	hbase2 "github.com/little-bit-shy/go-xgz/pkg/dao/hbase"
	"github.com/little-bit-shy/go-xgz/pkg/database/hbase"
	"github.com/little-bit-shy/go-xgz/pkg/helper"
)

// NewHbase new hbase pools
func NewHbase() (p *hbase2.Pool, cf func(), err error) {
	var (
		log        hbase2.Log
		cfg        hbase.Config
		poolConfig pool.Config
		ct         paladin.TOML
		zkCfg      hbase.ZKConfig
	)
	err = paladin.Get("hbase.toml").Unmarshal(&ct)
	helper.Panic(err)
	err = config.Env(&ct, &cfg, "Client")
	helper.Panic(err)
	err = config.Env(&ct, &zkCfg, "Zookeeper")
	helper.Panic(err)
	err = config.Env(&ct, &log, "Logrus")
	helper.Panic(err)
	cfg.Zookeeper = &zkCfg
	err = config.Env(&ct, &poolConfig, "Pool")
	helper.Panic(err)

	return hbase2.NewHbase(log, cfg, poolConfig, zkCfg)
}

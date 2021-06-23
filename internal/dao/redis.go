package dao

import (
	redis2 "github.com/little-bit-shy/go-xgz/pkg/cache/redis"
	"github.com/little-bit-shy/go-xgz/pkg/config"
	"github.com/little-bit-shy/go-xgz/pkg/dao/redis"
	"github.com/little-bit-shy/go-xgz/pkg/helper"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
)

func NewRedis() (r *redis.Redis, cf func(), err error) {
	var (
		cfg redis2.Config
		ct  paladin.Map
	)
	err = paladin.Get("redis.toml").Unmarshal(&ct)
	helper.Panic(err)
	err = config.Env(&ct, &cfg, "Client")
	helper.Panic(err)
	r = redis.New(&cfg)
	cf = func() { r.Close() }
	return
}

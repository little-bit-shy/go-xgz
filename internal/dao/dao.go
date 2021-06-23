package dao

import (
	"context"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/go-kratos/kratos/pkg/time"
	"github.com/little-bit-shy/go-xgz/internal/dao/client/api"
	"github.com/little-bit-shy/go-xgz/pkg/config"
	db2 "github.com/little-bit-shy/go-xgz/pkg/dao/db"
	"github.com/little-bit-shy/go-xgz/pkg/dao/es"
	"github.com/little-bit-shy/go-xgz/pkg/dao/hbase"
	"github.com/little-bit-shy/go-xgz/pkg/dao/redis"
	"github.com/little-bit-shy/go-xgz/pkg/helper"
	"time"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New, NewDB, NewRedis, NewHbase, NewEs, api.NewPool, NewJrpcPool)

//go:generate kratos tool genbts

type daoM interface {
	Close()
	Ping(ctx context.Context) (err error)
}

// Dao dao interface
type Dao struct {
	Db         *db2.Db
	Redis      *redis.Redis
	Hbase      *hbase.Pool
	Es         *es.Pool
	Jrpc       *Jrpc
	Cache      *fanout.Fanout
	Api        *api.Pool
	DemoExpire int32
	daoM
}

// New new a dao and return.
func New(r *redis.Redis, db *db2.Db, hbase *hbase.Pool, es *es.Pool, api *api.Pool, jrpc *Jrpc) (d *Dao, cf func(), err error) {
	return newDao(r, db, hbase, es, api, jrpc)
}

func newDao(r *redis.Redis, db *db2.Db, hbase *hbase.Pool, es *es.Pool, api *api.Pool, jrpc *Jrpc) (d *Dao, cf func(), err error) {
	var ct paladin.TOML
	var cfg struct {
		Expire xtime.Duration
	}
	err = paladin.Get("application.toml").Unmarshal(&ct)
	helper.Panic(err)
	err = config.Env(&ct, &cfg, "App")
	helper.Panic(err)
	d = &Dao{
		Db:         db,
		Redis:      r,
		Hbase:      hbase,
		Es:         es,
		Jrpc:       jrpc,
		Cache:      fanout.New("cache"),
		Api:        api,
		DemoExpire: int32(time.Duration(cfg.Expire) / time.Second),
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.Cache.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return nil
}

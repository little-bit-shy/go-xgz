package http

import (
	"fmt"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/ratelimit/bbr"
	pb "github.com/little-bit-shy/go-xgz/api"
	"github.com/little-bit-shy/go-xgz/pkg/config"
	"github.com/little-bit-shy/go-xgz/pkg/helper"
)

var svc pb.AppServer

type keepConfig struct {
	Keep string
}

// New new a bm server.
func New(s pb.AppServer) (engine *bm.Engine, err error) {
	var (
		cfg  bm.ServerConfig
		bbr  bbr.Config
		keep keepConfig
		ct   paladin.TOML
	)
	err = paladin.Get("http.toml").Unmarshal(&ct)
	helper.Panic(err)
	err = config.Env(&ct, &cfg, "Server")
	helper.Panic(err)
	err = config.Env(&ct, &bbr, "Bbr")
	helper.Panic(err)
	err = config.Env(&ct, &keep, "Keep")
	helper.Panic(err)
	svc = s
	engine = bm.DefaultServer(&cfg)
	limiter := bm.NewRateLimiter(&bbr)
	engine.Use(limiter.Limit())
	engine.Use(keepAlive(&keep))
	cors := bm.CORS([]string{"null"})
	engine.Use(cors)
	pb.RegisterAppBMServer(engine, s)
	//initRouter(engine)
	err = engine.Start()
	return
}

// keepAlive keep long connect
func keepAlive(cfg *keepConfig) bm.HandlerFunc {
	return func(c *bm.Context) {
		header := c.Writer.Header()
		if cfg.Keep == "also" {
			header.Set("Connection", "keep-alive")
		} else {
			header.Set("Keep-Alive", fmt.Sprintf("timeout=%s", cfg.Keep))
		}
		c.Next()
	}
}

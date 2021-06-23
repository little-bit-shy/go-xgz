package grpc

import (
	"github.com/go-kratos/kratos/pkg/ratelimit/bbr"
	pb "github.com/little-bit-shy/go-xgz/api"
	"github.com/little-bit-shy/go-xgz/pkg/config"
	"github.com/little-bit-shy/go-xgz/pkg/helper"
	"github.com/little-bit-shy/go-xgz/pkg/rpc/warden"
	"github.com/little-bit-shy/go-xgz/pkg/rpc/warden/ratelimiter"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
)

// New new a grpc server.
func New(svc pb.AppServer) (ws *warden.Server, err error) {
	var (
		server warden.ServerConfig
		bbr    bbr.Config
		ct     paladin.TOML
	)
	err = paladin.Get("grpc.toml").Unmarshal(&ct)
	helper.Panic(err)
	err = config.Env(&ct, &server, "Server")
	helper.Panic(err)
	err = config.Env(&ct, &bbr, "Bbr")
	helper.Panic(err)
	ws = warden.NewServer(&server)
	limiter := ratelimiter.New(&bbr)
	ws.Use(limiter.Limit())
	pb.RegisterAppServer(ws.Server(), svc)
	ws, err = ws.Start()
	return
}

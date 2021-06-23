package main

import (
	"flag"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/little-bit-shy/go-xgz/internal/di"
	"github.com/little-bit-shy/go-xgz/pkg/config"
	"github.com/little-bit-shy/go-xgz/pkg/helper"
	"github.com/little-bit-shy/go-xgz/pkg/language"

	"github.com/little-bit-shy/go-xgz/pkg/dao/trace/zipkin"

	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var (
		err          error
		logConfig    log.Config
		zipkinConfig zipkin.Config
		ct           paladin.TOML
		cct          paladin.TOML
	)
	flag.Parse()
	paladin.Init()
	language.RegisterCode()
	err = paladin.Get("log.toml").Unmarshal(&ct)
	helper.Panic(err)
	err = config.Env(&ct, &logConfig, "Server")
	helper.Panic(err)
	log.Init(&logConfig)
	defer log.Close()

	err = paladin.Get("trace.toml").Unmarshal(&cct)
	helper.Panic(err)
	err = config.Env(&cct, &zipkinConfig, "Server")
	helper.Panic(err)
	zipkin.Init(&zipkinConfig)

	log.Info("app start")
	_, closeFunc, err := di.InitApp()
	helper.Panic(err)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("app exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

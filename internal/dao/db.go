package dao

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/little-bit-shy/go-xgz/pkg/config"
	db2 "github.com/little-bit-shy/go-xgz/pkg/dao/db"
	"github.com/little-bit-shy/go-xgz/pkg/database/sql"
	"github.com/little-bit-shy/go-xgz/pkg/helper"
	"strings"
)

var _key = "{db}"

type Cfg struct {
	Db string
}

func NewDB() (db *db2.Db, cf func(), err error) {
	var (
		c   Cfg
		cfg sql.Config
		ct  paladin.TOML
	)
	err = paladin.Get("db.toml").Unmarshal(&ct)
	helper.Panic(err)
	err = config.Env(&ct, &c, "Client")
	helper.Panic(err)
	err = config.Env(&ct, &cfg, "Client")
	helper.Panic(err)
	cfg.DSN = strings.Replace(cfg.DSN, _key, c.Db, -1)
	for k, v := range cfg.ReadDSN {
		cfg.ReadDSN[k] = strings.Replace(v, _key, c.Db, -1)
	}
	db = db2.New(&cfg)
	cf = func() { _ = db.Close() }
	return
}

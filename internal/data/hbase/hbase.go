package hbase

import (
	"context"
	"github.com/little-bit-shy/go-xgz/internal/dao"
	"github.com/little-bit-shy/go-xgz/pkg/database/hbase"
	"github.com/tsuna/gohbase/hrpc"
)

// GetSomething get test
func GetSomething(d *dao.Dao, ctx context.Context) (result *hrpc.Result, err error) {
	var client *hbase.Client
	var cf func()
	if client, cf, err = d.Hbase.G(ctx); err != nil {
		return
	}
	defer cf()
	table := []byte("hdm:gofish_words_zh-cn")
	key := []byte("1")
	if result, err = client.Get(ctx, table, key); err != nil {
		return
	}
	return
}

// PutSomething put test
func PutSomething(d *dao.Dao, ctx context.Context) (result *hrpc.Result, err error) {
	var client *hbase.Client
	var cf func()
	if client, cf, err = d.Hbase.G(ctx); err != nil {
		return
	}
	defer cf()
	table := "hdm:gofish_words_zh-cn"
	key := "1"
	if result, err = client.PutStr(ctx, table, key, map[string]map[string][]byte{
		"info": {
			"text": []byte("2333"),
		},
	}); err != nil {
		return
	}
	return
}

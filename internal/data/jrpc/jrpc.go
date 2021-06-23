package jrpc

import (
	"context"
	"github.com/little-bit-shy/go-xgz/internal/dao"
	"github.com/little-bit-shy/go-xgz/pkg/dao/jsonRpc"
)

// Dict jrpc call to dict
func Dict(d *dao.Dao, ctx context.Context) (res jsonRpc.Result, err error) {
	var (
		con     *jsonRpc.JsonRpc
		cf      func()
		resoult jsonRpc.Result
	)
	if con, cf, err = d.Jrpc.Dict.Jp.G(ctx); err != nil {
		return
	}
	defer cf()
	if resoult, err = con.Call(ctx, "dict", []interface{}{
		"say hello",
		"en",
		[]string{
			"en",
			"zh-cn",
		},
	}); err != nil {
		return
	}
	res = resoult
	return
}

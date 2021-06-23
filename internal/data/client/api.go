package client

import (
	"context"
	"github.com/little-bit-shy/go-xgz/internal/dao"
	"github.com/little-bit-shy/go-xgz/pkg/api"
)

// CallSelf call self
func CallSelf(d *dao.Dao, ctx context.Context, name string) (HelloResp *api.HelloResp, err error) {
	var c *api.C
	var cf func()
	if c, cf, err = d.Api.G(ctx); err != nil {
		return
	}
	defer cf()
	if HelloResp, err = c.New.CallSelf(ctx, &api.HelloReq{
		Name: name,
	}); err != nil {
		return
	}
	return
}

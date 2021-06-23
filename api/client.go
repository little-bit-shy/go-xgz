package api

import (
	"context"
	"github.com/little-bit-shy/go-xgz/pkg/rpc/warden"
	"google.golang.org/grpc"
)

type Config struct {
	AppId string
}

type C struct {
	c   Config
	cfg *warden.ClientConfig
	New AppClient
	Err error
}

//New new client
func New(appid string, cfg *warden.ClientConfig, opts ...grpc.DialOption) *C {
	client := &C{Config{AppId: appid}, cfg, nil, nil}
	client.New, client.Err = newClient(client, opts...)
	return client
}

//NewClient new grpc client
func newClient(c *C, opts ...grpc.DialOption) (ac AppClient, err error) {
	client := warden.NewClient(c.cfg, opts...)
	cc, err := client.Dial(context.Background(), c.c.AppId)
	if err != nil {
		return nil, err
	}
	ac = NewAppClient(cc)
	return ac, nil
}

//Close close grpc client
func (c *C) Close() error {
	return nil
}

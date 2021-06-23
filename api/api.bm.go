// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

package api

import (
	"context"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"
)
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathAppPing = "/ping"
var PathAppSay = "/say"
var PathAppCallSelf = "/call_self"

// AppBMServer is the server API for App service.
// 服务方法
type AppBMServer interface {
	// ping
	Ping(ctx context.Context, req *google_protobuf1.Empty) (resp *google_protobuf1.Empty, err error)

	// say
	Say(ctx context.Context, req *HelloReq) (resp *SayResp, err error)

	// call_self
	CallSelf(ctx context.Context, req *HelloReq) (resp *HelloResp, err error)
}

var AppSvc AppBMServer

func appPing(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := AppSvc.Ping(c, p)
	c.JSON(resp, err)
}

func appSay(c *bm.Context) {
	p := new(HelloReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := AppSvc.Say(c, p)
	c.JSON(resp, err)
}

func appCallSelf(c *bm.Context) {
	p := new(HelloReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := AppSvc.CallSelf(c, p)
	c.JSON(resp, err)
}

// RegisterAppBMServer Register the blademaster route
func RegisterAppBMServer(e *bm.Engine, server AppBMServer) {
	AppSvc = server
	e.GET("/ping", appPing)
	e.GET("/say", appSay)
	e.GET("/call_self", appCallSelf)
}

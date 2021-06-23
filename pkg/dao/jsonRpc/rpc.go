package jsonRpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/pkg/container/pool"
	"github.com/go-kratos/kratos/pkg/net/trace"
	xtime "github.com/go-kratos/kratos/pkg/time"
	"io"
	"net"
	"time"
)

const (
	version = "1.0"
	_family = "jrpc_client"
	eof     = "\r\n\r\n"
)

type Result map[string]interface{}

type Config struct {
	Network   string
	Address   string
	Interface string
	Timeout   xtime.Duration
	Pool      *pool.Config
}

type JsonRpc struct {
	cfg    *Config
	client net.Conn
	io.Closer
}

// New new json rpc
func New(config *Config) (client *JsonRpc, err error) {
	var con net.Conn
	if con, err = net.Dial(config.Network, config.Address); err != nil {
		return
	}
	client = &JsonRpc{
		cfg:    config,
		client: con,
	}
	return
}

// Call call some method
func (j *JsonRpc) Call(ctx context.Context, method string, param []interface{}) (res Result, err error) {
	do := fmt.Sprintf("%v::%v::%v", version, j.cfg.Interface, method)
	var paramByte []byte
	if paramByte, err = json.Marshal(param); err == nil {
		if t, ok := trace.FromContext(ctx); ok {
			var internalTags []trace.Tag
			internalTags = append(internalTags, trace.TagString(trace.TagComponent, "rpc/jrpc"))
			internalTags = append(internalTags, trace.TagString(trace.TagPeerService, "rpc"))
			internalTags = append(internalTags, trace.TagString(trace.TagSpanKind, "client"))

			t = t.Fork(_family, "Rpc:"+method)
			t.SetTag(trace.String(trace.TagAddress, j.cfg.Address),
				trace.String(trace.TagComment, fmt.Sprintf("method(%s) args(%+v)", do, string(paramByte))))
			defer t.Finish(&err)
		}
	}

	if _, err = j.client.Write([]byte("")); err != nil {
		if err = j.Reconnect(); err != nil {
			return
		}
	}

	if err = j.client.SetDeadline(time.Now().Add(time.Duration(j.cfg.Timeout))); err != nil {
		return
	}

	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  do,
		"params":  param,
		"id":      "",
		"ext":     []string{},
	}
	var dataJson []byte
	if dataJson, err = json.Marshal(data); err != nil {
		return
	}
	if _, err = j.client.Write([]byte(dataJson)); err != nil {
		return
	}
	if _, err = j.client.Write([]byte(eof)); err != nil {
		return
	}

	var message []byte
	for {
		var msg [128]byte
		var length int
		if length, err = j.client.Read(msg[:]); err != nil {
			return
		}
		message = append(message, msg[:length]...)
		if length < 128 {
			break
		}
	}

	res = map[string]interface{}{}
	if err = json.Unmarshal(message, &res); err != nil {
		return
	}
	return
}

// Reconnect close than connect
func (j *JsonRpc) Reconnect() (err error) {
	_ = j.Close()
	var con net.Conn
	if con, err = net.Dial(j.cfg.Network, j.cfg.Address); err != nil {
		return
	}
	j.client = con
	return
}

// Close close this connect
func (j *JsonRpc) Close() error {
	return j.client.Close()
}

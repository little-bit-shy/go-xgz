package redis

import (
	"context"
	"github.com/little-bit-shy/go-xgz/pkg/cache/redis"
	"io"
)

type Redis struct {
	Cfg     *redis.Config
	Connect *redis.Redis
	io.Closer
}

// New new redis connect
func New(cfg *redis.Config) *Redis {
	db := redis.NewRedis(cfg)
	return &Redis{
		Cfg:     cfg,
		Connect: db,
	}
}

// Do do something
func (r *Redis) Do(ctx context.Context, commandName string, args ...interface{}) (reply interface{}, err error) {
	reply, err = r.Connect.Do(ctx, commandName, args...)
	return
}

func (r *Redis) Close() error {
	_ = r.Connect.Close()
	return nil
}

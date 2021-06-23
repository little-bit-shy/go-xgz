package redis

import (
	"context"
	"github.com/little-bit-shy/go-xgz/internal/dao"
)

// GetTest get test
func GetTest(d *dao.Dao, ctx context.Context) (reply interface{}, err error) {
	reply, err = d.Redis.Do(ctx, "GET", "test")
	return
}

// SetTest set test
func SetTest(d *dao.Dao, ctx context.Context) (reply interface{}, err error) {
	reply, err = d.Redis.Do(ctx, "SETEX", "test", 60, "233333")
	return
}

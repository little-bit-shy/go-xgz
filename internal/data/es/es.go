package es

import (
	"context"
	"github.com/little-bit-shy/go-xgz/internal/dao"
	"github.com/little-bit-shy/go-xgz/pkg/elastic"
)

// GetHealth get es health
func GetHealth(d *dao.Dao, ctx context.Context) (response *elastic.ClusterHealthResponse, err error) {
	var client *elastic.Client
	var cf func()
	client, cf, err = d.Es.G(ctx)
	if err != nil {
		return
	}
	defer cf()
	health := client.ClusterHealth()
	if response, err = health.Do(ctx); err != nil {
		return
	}
	return
}

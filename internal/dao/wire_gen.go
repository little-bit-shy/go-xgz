// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package dao

import (
	"github.com/little-bit-shy/go-xgz/internal/dao/client/api"
)

// Injectors from wire.go:

//go:generate kratos tool wire
func NewDao() (*Dao, func(), error) {
	redis, cleanup, err := NewRedis()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup2, err := NewDB()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	pool, cleanup3, err := NewHbase()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	esPool, cleanup4, err := NewEs()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	apiPool, cleanup5, err := api.NewPool()
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	jrpc, cleanup6, err := NewJrpcPool()
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dao, cleanup7, err := New(redis, db, pool, esPool, apiPool, jrpc)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return dao, func() {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

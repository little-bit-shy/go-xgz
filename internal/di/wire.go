// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"
	"github.com/little-bit-shy/go-xgz/internal/cli"
	"github.com/little-bit-shy/go-xgz/internal/dao"
	"github.com/little-bit-shy/go-xgz/internal/server/grpc"
	"github.com/little-bit-shy/go-xgz/internal/server/http"
	"github.com/little-bit-shy/go-xgz/internal/service"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, cli.Provider, service.Provider, http.New, grpc.New, NewApp))
}

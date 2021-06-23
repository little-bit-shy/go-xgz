// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package dao

import (
	"github.com/google/wire"
)

//go:generate kratos tool wire
func NewDao() (*Dao, func(), error) {
	panic(wire.Build(Provider))
}

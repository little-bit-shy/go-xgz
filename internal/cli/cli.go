package cli

import (
	"github.com/google/wire"
	"github.com/little-bit-shy/go-xgz/internal/dao"
	"github.com/little-bit-shy/go-xgz/internal/job"
	"github.com/little-bit-shy/go-xgz/pkg/cli"
	job2 "github.com/little-bit-shy/go-xgz/pkg/cli/job"
)

var Provider = wire.NewSet(New)

type Cfg struct {
}

// Service service.
type Cli struct {
	Dao *dao.Dao
}

// New new a service and return.
func New(d *dao.Dao) (c *Cli, cf func(), err error) {
	c = &Cli{
		Dao: d,
	}
	cf = c.Close

	var cl *cli.Cli
	cl, err = cli.RegisterAppCliServer()
	cl.Add("*/1 * * * * ?", job2.New("test", job.Test))
	cl.Start()

	return
}

// Close close the resource.
func (c *Cli) Close() {
}

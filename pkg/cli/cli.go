package cli

import (
	"github.com/robfig/cron/v3"
)

type Cli struct {
	cron *cron.Cron
}

// RegisterAppBMServer Register the cmd method
func RegisterAppCliServer() (cl *Cli, err error) {
	cl = &Cli{
		cron: cron.New(cron.WithSeconds()),
	}
	return
}

// Start start run
func (c *Cli) Start() {
	c.cron.Start()
}

func (c *Cli) Add(spec string, cmd cron.Job) (entryId cron.EntryID, err error) {
	entryId, err = c.cron.AddJob(spec, cmd)
	return
}

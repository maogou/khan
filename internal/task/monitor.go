package task

import (
	"smallBot/internal/sdk/gewe"

	"github.com/robfig/cron/v3"
)

type Monitor struct {
	Name    string
	sdk     *gewe.Gewe
	crontab *cron.Cron
}

func NewMonitor(name string, sdk *gewe.Gewe, crontab *cron.Cron) *Monitor {
	return &Monitor{
		Name:    name,
		sdk:     sdk,
		crontab: crontab,
	}
}

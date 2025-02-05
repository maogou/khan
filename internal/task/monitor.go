package task

import (
	"smallBot/internal/sdk/khan"

	"github.com/robfig/cron/v3"
)

type Monitor struct {
	Name    string
	sdk     *khan.Khan
	crontab *cron.Cron
}

func NewMonitor(name string, sdk *khan.Khan, crontab *cron.Cron) *Monitor {
	return &Monitor{
		Name:    name,
		sdk:     sdk,
		crontab: crontab,
	}
}

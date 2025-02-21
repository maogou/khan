package task

func (m *Monitor) Run(appIds []string) {
	m.restart(appIds)
	m.logrotate()

	m.crontab.Start()
}

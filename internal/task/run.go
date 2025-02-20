package task

func (m *Monitor) Run(appIds []string) {
	m.restart(appIds)

	m.crontab.Start()
}

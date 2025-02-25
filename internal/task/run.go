package task

func (m *Monitor) Run(appIds []string) {
	//m.restart(appIds)
	//m.forward(appIds)
	m.logrotate()

	m.crontab.Start()
}

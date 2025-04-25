package task

func (m *Monitor) Run() {
	m.restart()
	m.forward()
	m.forceLogout()
	m.logrotate()

	m.crontab.Start()
}

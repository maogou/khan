package task

func (m *Monitor) Run() {
	m.restart()
	m.forward()
	m.logrotate()

	m.crontab.Start()
}

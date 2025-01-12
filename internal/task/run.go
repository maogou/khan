package task

func (m *Monitor) Run() {
	m.restart()

	m.crontab.Start()
}

package login

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Confirm interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
}

type confirm struct {
	choices  []string
	cursor   int
	selected string
}

var _ Confirm = (*confirm)(nil)

func NewConfirm() Confirm {
	return &confirm{
		choices: []string{
			"① 我已在手机上确认登录,稍后你的手机上将会显示iPad微信已登录",
			"② 我已经在手机上取消登录",
		},
	}
}

func (c confirm) Init() tea.Cmd {
	return nil
}

func (c confirm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return c, tea.Quit

		case "up", "k":
			if c.cursor > 0 {
				c.cursor--
			}

		case "down", "j":
			if c.cursor < len(c.choices)-1 {
				c.cursor++
			}

		case "enter", " ":
			c.selected = c.choices[c.cursor]
			return c, tea.Quit
		}

	}

	return c, nil
}

func (c confirm) View() string {
	s := "请选择你要执行的操作:\n\n"

	for i, choice := range c.choices {
		cursor := " "
		if c.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\n如需退出操作,请按键盘上的 q 键退出操作.\n"

	return s
}

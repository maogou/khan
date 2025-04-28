package login

import (
	"maogou/khan/internal/sdk/khan"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)

func Login(sdk *khan.Khan) *cli.Command {
	return &cli.Command{
		Name:  "login",
		Usage: "快捷登录(危险:不是首次登录,请务必传递appId参数)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "appId",
				Value:    "",
				Usage:    "设备id(警告:首次登录时为空,掉线或者重新登录必须传递上次登录的appId,否则会触发风控)",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			p := tea.NewProgram(NewConfirm())
			quick := NewQuickLogin(sdk, p)
			appId := cCtx.String("appId")
			quick.Start(appId)

			return quick.err
		},
	}
}

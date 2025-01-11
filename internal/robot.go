package internal

import (
	"os"
	"smallBot/internal/command/serve"
	"smallBot/internal/command/version"
	"smallBot/internal/constant"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"github.com/rs/zerolog"
)

func init() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05.000"})
}

func NewRobotCommand() *cli.App {

	robot := &cli.App{
		Name:                  "robot",
		Usage:                 "超简单的微信机器人",
		Copyright:             "(c) 2025 kinyou_xy@foxmail.com",
		Version:               constant.VERSION,
		CustomAppHelpTemplate: constant.LOGO + cli.AppHelpTemplate,

		Commands: []*cli.Command{
			serve.Start(),
			version.Info(),
		},
	}

	return robot
}

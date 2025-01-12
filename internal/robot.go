package internal

import (
	"os"
	"path/filepath"
	"smallBot/internal/command/demo"
	"smallBot/internal/command/serve"
	"smallBot/internal/command/version"
	"smallBot/internal/constant"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"github.com/rs/zerolog"
)

func init() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
	log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05.000"})
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
			demo.Demo(),
		},
	}

	return robot
}

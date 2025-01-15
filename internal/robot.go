package internal

import (
	"os"
	"path/filepath"
	"smallBot/internal/command/auth"
	"smallBot/internal/command/demo"
	"smallBot/internal/command/license"
	"smallBot/internal/command/login"
	"smallBot/internal/command/serve"
	"smallBot/internal/command/upgrade"
	"smallBot/internal/command/version"
	"smallBot/internal/config"
	"smallBot/internal/constant"
	"smallBot/internal/sdk/gewe"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"github.com/rs/zerolog"
)

var (
	conf     config.Config
	sdk      *gewe.Gewe
	client   = resty.New()
	validate = validator.New()
)

func init() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
	log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05.000"})
}

func NewRobotCommand() *cli.App {
	conf = config.MustLoadConfig("")
	sdk = gewe.NewGeweSdk(&conf, client, validate)

	robot := &cli.App{
		Name:      "bolt",
		Usage:     "超简单超稳定的机器人",
		Copyright: "(c) 2025 kinyou_xy@foxmail.com",
		Version:   constant.VERSION,

		Commands: []*cli.Command{
			serve.Start(conf, sdk),
			version.Info(),
			auth.Verify(),
			login.Login(conf, sdk),
			upgrade.Upgrade(),
			license.Create(),
			demo.Demo(),
		},
	}

	return robot
}

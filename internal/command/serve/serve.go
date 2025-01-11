package serve

import (
	"smallBot/internal/config"
	"smallBot/internal/sdk/gewe"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

var (
	conf config.Config
	sdk  *gewe.Gewe
)

func Start() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "启动机器人服务",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: config.DefaultConfigPath,
				Usage: "项目配置文件",
			},
		},
		Before: func(cCtx *cli.Context) error {
			log.Info().Msg("正在初始化机器人配置,请稍等...")
			conf = config.MustLoadConfig(cCtx.String("config"))
			sdk = gewe.NewGeweSdk(&conf, resty.New(), validator.New())

			return nil
		},
		Action: func(cCtx *cli.Context) error {
			log.Info().Msg("开始启动机器人服务...")

			if err := run(conf, sdk); err != nil {
				log.Error().Err(err).Msg("启动机器人服务失败")
				return err
			}
			return nil
		},
		After: func(cCtx *cli.Context) error {
			log.Info().Msg("机器人服务已关闭")
			return nil
		},
	}
}

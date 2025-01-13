package serve

import (
	"smallBot/internal/config"
	"smallBot/internal/sdk/gewe"
	"smallBot/internal/task"

	"github.com/robfig/cron/v3"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

var (
	monitor *task.Monitor
	crontab = cron.New()
)

func Start(conf config.Config, sdk *gewe.Gewe) *cli.Command {
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
			log.Info().Msg("初始化长连接监控...")
			monitor = task.NewMonitor("长连接监控", sdk, crontab)
			return nil
		},

		Action: func(cCtx *cli.Context) error {
			log.Info().Msg("开启定时任务...")
			monitor.Run()
			log.Info().Msg("开始启动机器人服务...")

			if err := run(conf, sdk); err != nil {
				log.Error().Err(err).Msg("启动机器人服务失败")
				return err
			}
			return nil
		},
		After: func(cCtx *cli.Context) error {
			crontab.Stop()
			log.Info().Msg("机器人服务已关闭")
			return nil
		},
	}
}

package serve

import (
	"errors"
	"os"
	"path/filepath"
	"smallBot/internal/config"
	"smallBot/internal/pkg/license"
	"smallBot/internal/sdk/gewe"
	"smallBot/internal/task"
	"strings"

	"github.com/robfig/cron/v3"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

var (
	monitor *task.Monitor
	crontab = cron.New()
	lic     = &license.License{}
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
			if _, err := os.Stat(conf.Sdk.License); err != nil {
				return err
			}

			pKey := filepath.Base(conf.Sdk.License)
			pKey = strings.ReplaceAll(pKey, "37", "+")
			pKey = strings.ReplaceAll(pKey, "73", "/")

			nLic, err := license.Parse(pKey, conf.Sdk.License)
			if err != nil {
				log.Error().Err(err).Msg("许可证校验失败")
				return err
			}

			if nLic.Expired() {
				log.Error().Msg("许可证已过期")
				return errors.New("授权许可证已过期")
			}

			lic = nLic

			return nil
		},

		Action: func(cCtx *cli.Context) error {
			log.Info().Msg("开启定时任务...")
			monitor.Run()
			log.Info().Msg("开始启动机器人服务...")

			if err := run(conf, sdk, lic); err != nil {
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

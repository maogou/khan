package serve

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/config"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/license"
	"smallBot/internal/sdk/khan"
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
	p       v1.Permission
	appIds  = make([]string, 0)
)

func Start(conf config.Config, sdk *khan.Khan) *cli.Command {
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
			oPKey := pKey
			pKey = strings.NewReplacer(
				constant.License37, "+",
				constant.License73, "/",
				constant.License919, "=",
			).Replace(pKey)

			nLic, err := license.Parse(pKey, conf.Sdk.License)
			if err != nil {
				log.Error().Msg("许可证校验失败")
				return err
			}

			if nLic.Expired() {
				log.Error().Msg("许可证已过期")
				return errors.New("授权许可证已过期")
			}

			lic = nLic

			kl, _ := os.ReadFile(conf.Sdk.License)

			if err = sdk.Rdb().Set(cCtx.Context, constant.License, string(kl), 0).Err(); err != nil {
				log.Error().Msg("redis设置授权许可证失败")
				return err
			}

			if err = sdk.Rdb().Set(cCtx.Context, constant.LicenseKey, oPKey, 0).Err(); err != nil {
				log.Error().Msg("redis设置授权许可证key失败")
				return err
			}

			if err = json.Unmarshal(lic.Dat, &p); err != nil {
				log.Error().Err(err).Msg("json解析授权许可证失败")
				return err
			}

			appIds = p.AppId

			return nil
		},

		Action: func(cCtx *cli.Context) error {
			log.Info().Msg("开启定时任务...")
			monitor.Run(appIds)
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

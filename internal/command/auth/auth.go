package auth

import (
	"os"
	"path/filepath"
	"smallBot/internal/config"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/license"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/pterm/pterm"

	"github.com/urfave/cli/v2"
)

func Verify(conf config.Config) *cli.Command {
	return &cli.Command{
		Name:  "verify",
		Usage: "授权许可证验证",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "path",
				Value: conf.Sdk.License,
				Usage: "颁发的软件许可证绝对路径",
			},
		},
		Action: func(cCtx *cli.Context) error {
			path := cCtx.String("path")
			if _, err := os.Stat(path); len(path) > 0 && err != nil {
				return err
			}

			pKey := filepath.Base(path)
			pKey = strings.ReplaceAll(pKey, constant.License37, "+")
			pKey = strings.ReplaceAll(pKey, constant.License73, "/")
			pKey = strings.ReplaceAll(pKey, constant.License919, "=")

			lic, err := license.Parse(pKey, path)

			if err != nil {
				log.Error().Err(err).Msg("授权许可证验证失败")
				return err
			}

			if lic.Expired() {
				log.Error().Err(err).Msg("授权许可证已过期")
				return err
			}

			_ = printLicense(lic)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func printLicense(license *license.License) error {
	return pterm.DefaultTable.WithHasHeader().
		WithBoxed().
		WithData(
			pterm.TableData{
				{"配置项", "属性值"},
				{"微信号", license.Sub},
				{"发证单位", license.Iss},
				{"共享数量", strconv.FormatInt(int64(license.Lim), 10)},
				{"签发时间", license.Iat.Format("2006-01-02 15:04:05")},
				{"过期时间", license.Exp.Format("2006-01-02 15:04:05")},
			},
		).Render()
}

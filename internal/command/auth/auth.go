package auth

import (
	"encoding/json"
	"os"
	"path/filepath"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/config"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/help"
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

			pKey = strings.NewReplacer(
				constant.License37, "+",
				constant.License73, "/",
				constant.License919, "=",
			).Replace(pKey)

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

	var (
		p v1.Permission
	)
	if err := json.Unmarshal(license.Dat, &p); err != nil {
		log.Error().Err(err).Msg("json解析授权许可证失败")
		return err
	}

	names := help.TransferPermissionToName(p.Permission)

	nb, err := json.Marshal(names)
	if err != nil {
		log.Error().Err(err).Msg("json解析授权许可证(permission-名称)失败")
		return err
	}

	wb, err := json.Marshal(p.AppId)
	if err != nil {
		log.Error().Err(err).Msg("json解析授权许可证(wid)失败")
		return err
	}

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
				{"可用权限", string(nb)},
				{"共享微信", string(wb)},
			},
		).Render()
}

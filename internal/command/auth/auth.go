package auth

import (
	"context"
	"encoding/json"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/pkg/help"
	"maogou/khan/internal/pkg/license"
	"maogou/khan/internal/sdk/khan"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/pterm/pterm"

	"github.com/urfave/cli/v2"
)

func Verify(sdk *khan.Khan) *cli.Command {
	return &cli.Command{
		Name:  "verify",
		Usage: "授权许可证验证",
		Action: func(cCtx *cli.Context) error {
			lic, err := help.License(context.Background(), sdk.Rdb())
			if err != nil {
				log.Error().Err(err).Msg("授权许可证验证失败")
				return err
			}

			if err = printLicense(lic); err != nil {
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

	token, err := json.Marshal(p.Token)

	if err != nil {
		log.Error().Err(err).Msg("json解析授权许可证(token)失败")
		return err
	}

	return pterm.DefaultTable.WithHasHeader().
		WithBoxed().
		WithData(
			pterm.TableData{
				{"配置项", "属性值"},
				{"客户标识", license.Cus},
				{"发证单位", license.Iss},
				{"用户类型", license.Typ},
				{"共享数量", strconv.FormatInt(int64(license.Lim), 10)},
				{"签发时间", license.Iat.Format("2006-01-02 15:04:05")},
				{"过期时间", license.Exp.Format("2006-01-02 15:04:05")},
				{"可用权限", string(nb)},
				{"共享appid", string(wb)},
				{"token", string(token)},
			},
		).Render()
}

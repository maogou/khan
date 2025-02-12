package license

import (
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/help"
	"smallBot/internal/pkg/license"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

const licensePath = "licenses"

func Create() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "生成许可证",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "appid",
				Value:    "",
				Usage:    "应用的appid",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "account",
				Value:    "",
				Usage:    "登陆的账号",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "day",
				Value: "7",
				Usage: "许可证的可用天数",
			},
			&cli.StringFlag{
				Name:  "limit",
				Value: "1",
				Usage: "许可证的共享人数",
			},
			&cli.BoolFlag{
				Name:  "vip",
				Value: false,
				Usage: "vip用户享受高级功能",
			},
		},
		Action: func(cCtx *cli.Context) error {
			return do(cCtx)
		},
	}
}

func do(cCtx *cli.Context) error {
	out := cCtx.String("appid")
	current := time.Now().Local()
	vip := cCtx.Bool("vip")

	permission := lo.Ternary(vip, help.AccessVipPermission(), help.AccessPermission())

	pb, _ := help.TransferPermissionToByte(permission)

	days := lo.Ternary(cCtx.Int("day") > 0, cCtx.Int("day"), 7)
	l := license.License{
		Iss: "khan",
		Cus: cCtx.String("account"),
		Sub: out,
		Lim: 1,
		Iat: current,
		Exp: current.AddDate(0, 0, days),
		Dat: json.RawMessage(pb),
	}

	filePath, err := filepath.Abs(licensePath)
	if err != nil {
		log.Error().Err(err).Msg("获取路径失败")
		return err
	}

	if err = l.KeyGen(filePath, out); err != nil {
		log.Error().Err(err).Msg("生成失败")
		return err
	}

	log.Info().Msg("生成秘钥成功")

	pKByte, err := os.ReadFile(filePath + "/" + out + ".pub")
	if err != nil {
		log.Error().Err(err).Msg("获取公钥失败")
		return err
	}

	pKey := strings.ReplaceAll(string(pKByte), "+", constant.License37)
	pKey = strings.ReplaceAll(pKey, "/", constant.License73)
	pKey = strings.ReplaceAll(pKey, "=", constant.License919)

	if err = l.Create(filePath+"/"+out+".pri", filePath+"/"+pKey); err != nil {
		log.Error().Err(err).Msg("创建许可证失败")
		return err
	}

	log.Info().Msg("创建许可证成功")

	nl, err := l.Verify(filePath+"/"+out+".pub", filePath+"/"+pKey)

	if err != nil {
		log.Error().Err(err).Msg("验证失败")
		return err
	}

	nlByte, err := json.MarshalIndent(nl, "", "    ")

	if err != nil {
		log.Error().Err(err).Msg("json序列化失败")
		return err
	}

	fmt.Println(string(nlByte))

	return nil
}

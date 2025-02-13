package license

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/help"
	"smallBot/internal/pkg/license"
	"strings"
	"time"

	"github.com/samber/lo"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

const licensePath = "licenses"

var p v1.Permission

func Create() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "生成许可证",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "appid",
				Value:    cli.NewStringSlice(),
				Usage:    "应用的appid --appid=id1 --appid=id2",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "account",
				Value:    "",
				Usage:    "登陆的账号",
				Required: true,
			},
			&cli.IntFlag{
				Name:  "day",
				Value: 7,
				Usage: "许可证的可用天数",
			},
			&cli.IntFlag{
				Name:  "limit",
				Value: 1,
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
	appIds := cCtx.StringSlice("appid")

	firstAppId := appIds[0]
	current := time.Now().Local()
	vip := cCtx.Bool("vip")

	permission := lo.Ternary(vip, help.AccessVipPermission(), help.AccessPermission())

	p.Permission = permission
	p.Wid = appIds

	pb, _ := json.Marshal(&p)

	days := lo.Ternary(cCtx.Int("day") > 0, cCtx.Int("day"), 7)
	limit := lo.Ternary(cCtx.Int("limit") > 0, cCtx.Int("limit"), 1)
	l := license.License{
		Iss: "khan",
		Cus: cCtx.String("account"),
		Sub: firstAppId,
		Lim: limit,
		Iat: current,
		Exp: current.AddDate(0, 0, days),
		Dat: json.RawMessage(pb),
	}

	filePath, err := filepath.Abs(licensePath)
	if err != nil {
		log.Error().Err(err).Msg("获取路径失败")
		return err
	}

	if err = l.KeyGen(filePath, firstAppId); err != nil {
		log.Error().Err(err).Msg("生成失败")
		return err
	}

	log.Info().Msg("生成秘钥成功")

	pKByte, err := os.ReadFile(filePath + "/" + firstAppId + ".pub")
	if err != nil {
		log.Error().Err(err).Msg("获取公钥失败")
		return err
	}

	pKey := string(pKByte)

	pKey = strings.NewReplacer(
		"+", constant.License37,
		"/", constant.License73,
		"=", constant.License919,
	).Replace(pKey)

	if err = l.Create(filePath+"/"+firstAppId+".pri", filePath+"/"+pKey); err != nil {
		log.Error().Err(err).Msg("创建许可证失败")
		return err
	}

	log.Info().Msg("创建许可证成功")

	nl, err := l.Verify(filePath+"/"+firstAppId+".pub", filePath+"/"+pKey)

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

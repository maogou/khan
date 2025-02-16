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
	"smallBot/internal/sdk/khan"
	"strings"
	"time"

	"github.com/samber/lo"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

const licensePath = "licenses"

var p v1.Permission

func Create(sdk *khan.Khan) *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "生成许可证",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "appid",
				Value:    cli.NewStringSlice(),
				Usage:    "应用的appid --appid=id1 --appid=id2 支持批量",
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
			&cli.StringFlag{
				Name:     "account",
				Value:    "",
				Usage:    "授权的账号",
				Required: true,
			},
		},
		Action: func(cCtx *cli.Context) error {
			return do(cCtx, sdk)
		},
	}
}

func do(cCtx *cli.Context, sdk *khan.Khan) error {
	appIds := cCtx.StringSlice("appid")
	account := cCtx.String("account")
	current := time.Now().Local()
	vip := cCtx.Bool("vip")

	permission := lo.Ternary(vip, help.AccessVipPermission(), help.AccessPermission())
	days := lo.Ternary(cCtx.Int("day") > 0, cCtx.Int("day"), 7)
	limit := lo.Ternary(cCtx.Int("limit") > 0, cCtx.Int("limit"), 1)

	p.AppId = appIds
	p.Permission = permission
	p.Token = make(map[string]string)

	for _, appId := range appIds {
		encrypted, err := help.AesEncrypt(constant.AppName+appId, []byte(constant.AesWXidKey))
		if err != nil {
			log.Error().Err(err).Msg("token 加密失败")
			return err
		}
		p.Token[appId] = encrypted
	}

	pb, _ := json.Marshal(&p)

	l := license.License{
		Iss: constant.AppName,
		Cus: account,
		Sub: account,
		Typ: lo.Ternary(vip, "vip用户", "普通用户"),
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

	if err = l.KeyGen(filePath, account); err != nil {
		log.Error().Err(err).Msg("生成失败")
		return err
	}

	log.Info().Msg("生成秘钥成功")

	pKByte, err := os.ReadFile(filePath + "/" + account + ".pub")
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

	if err = l.Create(filePath+"/"+account+".pri", filePath+"/"+pKey); err != nil {
		log.Error().Err(err).Msg("创建许可证失败")
		return err
	}

	log.Info().Msg("创建许可证成功")

	nl, err := l.Verify(filePath+"/"+account+".pub", filePath+"/"+pKey)

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

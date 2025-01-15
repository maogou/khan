package demo

import (
	"encoding/json"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/license"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/urfave/cli/v2"
)

func Demo() *cli.Command {
	return &cli.Command{
		Name:  "demo",
		Usage: "测试khan的接口",
		Action: func(cCtx *cli.Context) error {
			return do(cCtx)
		},
	}
}

func do(cCtx *cli.Context) error {
	l := license.License{
		Iss: "khan",
		Sub: "微信号",
		Lim: 1,
		Iat: time.Now().Local(),
		Exp: time.Now().Local().Add(time.Hour * 24 * 7),
	}

	if err := l.KeyGen("./khan"); err != nil {
		log.Error().Err(err).Msg("生成失败")
		return err
	}

	log.Info().Msg("生成成功")

	if err := l.Create("./khan.pri", "./khan.lic"); err != nil {
		log.Error().Err(err).Msg("创建许可证失败")
		return err
	}

	log.Info().Msg("创建许可证成功")

	nl, err := l.Verify(constant.PublicKey, "./khan.lic")

	if err != nil {
		log.Error().Err(err).Msg("验证失败")
		return err
	}

	nlByte, err := json.Marshal(nl)

	if err != nil {
		log.Error().Err(err).Msg("json序列化失败")
		return err
	}

	log.Info().Str("license", string(nlByte)).Msg("验证成功")

	return nil
}

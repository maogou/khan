package demo

import (
	"github.com/rs/zerolog/log"
	"smallBot/internal/sdk/khan"

	"github.com/urfave/cli/v2"
)

func Demo(sdk *khan.Khan) *cli.Command {
	return &cli.Command{
		Name:  "demo",
		Usage: "测试khan的接口",

		Action: func(cCtx *cli.Context) error {
			return do(cCtx, sdk)
		},
	}
}

func do(cCtx *cli.Context, sdk *khan.Khan) error {
	log.Info().Msg("i am demo!")

	_, err := sdk.Tpl().ReadFile("tpl/text.tpl")

	if err != nil {
		log.Error().Err(err).Msg("读取模板失败")
		return err
	}

	log.Info().Msg("读取模板成功")
	return nil
}

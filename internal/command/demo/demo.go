package demo

import (
	"smallBot/internal/sdk/khan"

	"github.com/rs/zerolog/log"

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
	log.Info().Msg("demo")
	return nil
}

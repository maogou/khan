package demo

import (
	v1 "smallBot/api/gewe/v1"
	"smallBot/internal/config"
	"smallBot/internal/sdk/gewe"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"

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
	conf := config.MustLoadConfig("")
	client := resty.New()
	validate := validator.New()
	sdk := gewe.NewGeweSdk(&conf, client, validate)

	resp, err := sdk.PostText(
		cCtx.Context, v1.PostTextRequest{
			Appid: "wx_Ez4I3ZO1gKOG7h0IgfB8D",
			ToWxidList: []v1.ToWxid{
				{
					Content: "当前时间为:" + time.Now().Format(time.DateTime),
					ToWxid:  "xingmaogou",
					MsgType: 1,
				},
			},
		},
	)

	if err != nil {
		log.Info().Any("err", err).Msg("err")
		return err
	}

	log.Info().Any("resp", resp).Msg("resp")
	return nil
}

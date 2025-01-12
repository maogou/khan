package demo

import (
	"context"
	v1 "smallBot/api/gewe/v1"
	"smallBot/internal/config"
	"smallBot/internal/sdk/gewe"
	"time"

	"github.com/rs/xid"
	"github.com/rs/zerolog/log"

	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"

	"github.com/urfave/cli/v2"
)

func Demo() *cli.Command {
	return &cli.Command{
		Name:  "demo",
		Usage: "测试gewe的接口",
		Action: func(cCtx *cli.Context) error {
			return do(cCtx)
		},
	}
}

func do(cCtx *cli.Context) error {

	conf := config.MustLoadConfig(cCtx.String("config"))
	sdk := gewe.NewGeweSdk(&conf, resty.New(), validator.New())

	log.Info().Msg("开始获取appid...")

	ctx := context.WithValue(context.Background(), "qid", xid.New().String())

	sdkResp, err := sdk.CreateApp(
		ctx, v1.CreateAppRequest{
			Country:    "中国",
			DeviceName: "iPad",
			Model:      "iPad",
			SdkVer:     "8.0.48",
		},
	)

	if err != nil {
		log.Error().Err(err).Msg("获取appid失败")
		return err
	}

	if sdkResp.Ret != 0 {
		log.Info().Int("ret", sdkResp.Ret).Msg("sdkResp.Ret !=0")
		return err
	}

	log.Info().Str("appid", sdkResp.Data.AppId).Msg("获取appid成功")

	log.Info().Msg("开始获取二维码...")

	lqcResp, err := sdk.LoginQrCode(
		ctx, v1.LoginQrCodeRequest{
			AppId: sdkResp.Data.AppId,
		},
	)

	if err != nil {
		log.Error().Err(err).Msg("获取二维码失败")
		return err
	}

	if lqcResp.Ret != 0 {
		log.Info().Int("ret", lqcResp.Ret).Msg("lqcResp.Ret !=0")
		return err
	}

	log.Info().Str("base64", lqcResp.Data.Base64).Msg("二维码")
	log.Info().Str("appid", sdkResp.Data.AppId).Msg("appid")
	log.Info().Str("nkey", lqcResp.Data.NKey).Msg("nkey")
	log.Info().Str("uuid", lqcResp.Data.Uuid).Msg("uuid")
	log.Info().Msg("确认登陆开始....")

	time.Sleep(60 * time.Second)

	clqrResp, err := sdk.CheckLoginQrCode(
		ctx, v1.CheckLoginQrCodeRequest{
			AppId: sdkResp.Data.AppId,
			NKey:  lqcResp.Data.NKey,
			Uuid:  lqcResp.Data.Uuid,
		},
	)

	if err != nil {
		log.Error().Err(err).Msg("确认登陆失败")
		return err
	}

	if clqrResp.Ret != 0 {
		log.Info().Int("ret", clqrResp.Ret).Msg("clqrResp.Ret !=0")
		return err
	}

	log.Info().Any("clqrResp", clqrResp).Msg("确认登陆成功")

	log.Info().Msg("开启长链服务并设置消息收集接口.....")

	loResp, err := sdk.LongOpen(
		ctx, v1.LongOpenRequest{
			AppId:      sdkResp.Data.AppId,
			CleanCache: true,
			Host:       "http://127.0.0.1:8073/api/v1/collect",
			Timeout:    10,
		},
	)

	if err != nil {
		log.Error().Err(err).Msg("开启长链服务并设置消息收集接口失败")
		return err
	}

	if loResp.Ret != 0 {
		log.Info().Int("ret", loResp.Ret).Msg("loResp.Ret !=0")
		return err
	}

	log.Info().Msg("自动授权appid 开始....")

	saaResp, err := sdk.SecAutoAuth(
		ctx, v1.SecAutoAuthRequest{
			AppId:  sdkResp.Data.AppId,
			SdkVer: "8.0.48",
		},
	)

	if err != nil {
		log.Error().Err(err).Msg("自动授权appid 失败")
		return err
	}

	if saaResp.Ret != 0 {
		log.Info().Int("ret", saaResp.Ret).Msg("saaResp.Ret !=0")
		return err
	}

	log.Info().Any("saaResp", saaResp).Msg("自动授权appid 成功")

	return nil
}

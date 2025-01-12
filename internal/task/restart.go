package task

import (
	"context"
	v1 "smallBot/api/gewe/v1"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/help"

	"github.com/rs/zerolog/log"

	"github.com/rs/xid"
)

func (m *Monitor) restart() {
	appId := m.sdk.Config().Sdk.AppId

	log.Info().Str("appId", appId).Msg("开始监控重启长连接.....")

	m.crontab.AddFunc(
		"@every 120s", func() {
			ctx := context.WithValue(context.Background(), constant.QID, xid.New().String())
			log := help.GetQidLog(ctx)
			loResp, err := m.sdk.LongOpen(
				ctx, v1.LongOpenRequest{
					AppId:      appId,
					Timeout:    60,
					CleanCache: true,
					Host:       "http://127.0.0.1:8073/api/v1/collect",
				},
			)

			if err != nil {
				log.Error().Err(err).Msg("调用开启长连接失败")
			} else {
				log.Info().Any("loResp", loResp).Msg("调用长连接返回结果")

				if loResp.Ret != 0 {
					log.Info().Int("ret", loResp.Ret).Msg("loResp.Ret !=0")
				} else {
					log.Info().Msg("开启长连接成功")
				}
			}
		},
	)
}

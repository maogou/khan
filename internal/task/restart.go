package task

import (
	"context"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/tranfer"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/log"
	"strconv"

	"github.com/rs/xid"
)

func (m *Monitor) restart(appIds []string) {
	var tbhb tranfer.TranferBatchHeartBeatRequest

	for _, appId := range appIds {
		tbhb.List = append(
			tbhb.List, tranfer.TranferBatchHeartBeatItem{
				Appid: appId,
			},
		)
	}

	ctx := context.WithValue(context.Background(), constant.QID, xid.New().String())
	collectUrl := m.getCollectUrl(strconv.Itoa(m.sdk.Config().Port))

	log.C(ctx).Info().Strs("appId", appIds).Msg("开始监控重启长连接.....")

	m.crontab.AddFunc(
		"@every 60s", func() {
			ctx = context.WithValue(context.Background(), constant.QID, xid.New().String())

			for _, appId := range appIds {
				if longStatus, err := m.getLongConnectTime(ctx, appId); err == nil && longStatus == constant.WXLongOK {
					continue
				}

				if ok := m.reconnection(ctx, appId, collectUrl); ok {
					if err := m.setLongConnectTime(ctx, appId); err != nil {
						log.C(ctx).Error().Err(err).Str("appId", appId).Msg("设置长连接时间失败")
					} else {
						log.C(ctx).Info().Str("appId", appId).Msg("设置长连接时间成功")
					}
				}
			}
		},
	)
}

func (m *Monitor) getCollectUrl(port string) string {
	return "http://127.0.0.1:" + port + "/api/v1/collect"
}

func (m *Monitor) reconnection(ctx context.Context, appid, collectUrl string) bool {
	log.C(ctx).Warn().Str("appId", appid).Msg("💗心跳异常,开始监控重启长连接.....")
	loResp, lErr := m.sdk.LongOpen(
		ctx, v1.LongOpenRequest{
			AppId:      appid,
			Timeout:    60,
			CleanCache: true,
			Host:       collectUrl,
		},
	)

	if lErr != nil {
		log.C(ctx).Error().Err(lErr).Msg("调用开启长连接失败")
		return false
	}
	log.C(ctx).Info().Any("loResp", loResp).Msg("调用长连接返回结果")
	if loResp.Ret != 0 {
		log.C(ctx).Info().Int("ret", loResp.Ret).Msg("loResp.Ret !=0")
		return false
	}

	log.C(ctx).Info().Msg("开启长连接成功")
	return true
}

func (m *Monitor) setLongConnectTime(ctx context.Context, appid string) error {
	return m.sdk.Rdb().SetEx(ctx, constant.WXLongCache+appid, constant.WXLongOK, constant.WXQrCodeExpire*3).Err()
}

func (m *Monitor) getLongConnectTime(ctx context.Context, appid string) (string, error) {
	return m.sdk.Rdb().Get(ctx, constant.WXLongCache+appid).Result()
}

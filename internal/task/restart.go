package task

import (
	"context"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/help"
	"maogou/khan/internal/pkg/log"
	"strconv"
	"time"

	"github.com/rs/xid"
)

func (m *Monitor) restart() {

	ctx := context.WithValue(context.Background(), constant.QID, xid.New().String())
	collectUrl := m.getCollectUrl(strconv.Itoa(m.sdk.Config().Port))

	log.C(ctx).Info().Msg("开始监控重启长连接.....")

	m.crontab.AddFunc(
		"@every 60s", func() {
			ctx = context.WithValue(context.Background(), constant.QID, xid.New().String())
			var appIds []string
			p, err := help.Permission(ctx, m.sdk.Rdb())
			if err != nil {
				log.C(ctx).Error().Err(err).Msg("获取license失败")
				return
			}

			appIds = p.AppId

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
		if loResp.MsgErr == constant.WXLongAlreadyConnect {
			log.C(ctx).Info().Int("ret", loResp.Ret).Msg(loResp.MsgErr)
			return true
		}
		return false
	}

	log.C(ctx).Info().Msg("开启长连接成功")
	return true
}

func (m *Monitor) setLongConnectTime(ctx context.Context, appid string) error {
	return m.sdk.Rdb().SetEx(
		ctx, constant.WXLongCache+appid, constant.WXLongOK, constant.WXQrCodeExpire*3*time.Second,
	).Err()
}

func (m *Monitor) getLongConnectTime(ctx context.Context, appid string) (string, error) {
	return m.sdk.Rdb().Get(ctx, constant.WXLongCache+appid).Result()
}

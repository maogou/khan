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

			bhbResp, err := m.sdk.BatchHeartBeat(ctx, tbhb)
			if err != nil {
				log.C(ctx).Error().Err(err).Msg("调用批量心跳失败")
				return
			}

			if bhbResp.Ret != 0 {
				log.C(ctx).Error().Int("ret", bhbResp.Ret).Msg("bhbResp.Ret !=0 调用批量心跳失败")
				return
			}

			for _, item := range bhbResp.Data.List {

				switch v := item.ServiceInfo.(type) {
				case tranfer.ServiceInfo:
					var serviceInfo tranfer.ServiceInfo
					if serviceInfo.BaseResponse.Ret != 0 {
						m.reconnection(ctx, item.Appid, collectUrl)
					}
				case string:
					log.C(ctx).Error().Str("err_msg", v).Msg("💗心跳异常")
					if v == constant.WXOffline {
						m.reconnection(ctx, item.Appid, collectUrl)
					}
				default:
					log.C(ctx).Error().Any("other", v).Msg("其他情况，先监控暂不处理")
				}

			}
		},
	)
}

func (m *Monitor) getCollectUrl(port string) string {
	return "http://127.0.0.1:" + port + "/api/v1/collect"
}

func (m *Monitor) reconnection(ctx context.Context, appid, collectUrl string) {
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
	} else {
		log.C(ctx).Info().Any("loResp", loResp).Msg("调用长连接返回结果")
		if loResp.Ret != 0 {
			log.C(ctx).Info().Int("ret", loResp.Ret).Msg("loResp.Ret !=0")
		} else {
			log.C(ctx).Info().Msg("开启长连接成功")
		}
	}
}

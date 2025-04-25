package task

import (
	"context"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/help"
	"smallBot/internal/pkg/log"
	"smallBot/internal/task/forward/processing"

	"github.com/rs/xid"
	"github.com/sourcegraph/conc"
)

func (m *Monitor) forward() {
	m.crontab.AddFunc(
		"@every 35s", func() {
			var (
				wg     conc.WaitGroup
				appIds []string
			)
			chain := processing.NewForwardChain(m.sdk)
			ctx := context.WithValue(context.Background(), constant.QID, xid.New().String())

			log.C(ctx).Info().Msg("开始监控转发消息.....")

			p, err := help.Permission(ctx, m.sdk.Rdb())

			if err != nil {
				log.C(ctx).Error().Err(err).Msg("获取license失败")
				return
			}

			appIds = p.AppId

			for _, appId := range appIds {
				wg.Go(
					func() {
						ctx = context.WithValue(context.Background(), constant.QID, xid.New().String())

						if err = chain.Execute(ctx, appId); err != nil {
							log.C(ctx).Error().Err(err).Str("appId", appId).Msg("处理流程失败")
							return
						}
					},
				)
			}

			wg.Wait()
		},
	)
}

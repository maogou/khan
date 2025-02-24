package task

import (
	"context"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/log"
	"smallBot/internal/task/forward/processing"

	"github.com/rs/xid"
	"github.com/sourcegraph/conc"
)

func (m *Monitor) forward(appIds []string) {
	m.crontab.AddFunc(
		"@every 30s", func() {

			var wg conc.WaitGroup
			chain := processing.NewForwardChain(m.sdk)

			log.C(context.Background()).Info().Msg("开始监控转发消息.....")

			for _, appId := range appIds {
				wg.Go(
					func() {
						ctx := context.WithValue(context.Background(), constant.QID, xid.New().String())

						if err := chain.Execute(ctx, appId); err != nil {
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

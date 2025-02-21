package task

import (
	"context"
	"github.com/rs/xid"
	"os/exec"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/log"
)

func (m *Monitor) logrotate() {
	m.crontab.AddFunc(
		"0 2 * * *", func() {
			ctx := context.WithValue(context.Background(), constant.QID, xid.New().String())
			cmd := exec.CommandContext(ctx, "/usr/sbin/logrotate", "-f", "/etc/logrotate.conf")

			if err := cmd.Run(); err != nil {
				log.C(ctx).Error().Err(err).Msg("logrotate error")
			} else {
				log.C(ctx).Info().Msg("logrotate success")
			}
		},
	)
}

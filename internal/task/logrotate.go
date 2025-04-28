package task

import (
	"context"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/log"
	"os/exec"

	"github.com/rs/xid"
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

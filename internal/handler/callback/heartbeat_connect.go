package callback

import (
	"context"
	"smallBot/internal/pkg/help"
)

func (c *CallbackHandler) Heartbeat(ctx context.Context) {
	log := help.GetQidLog(ctx)
	log.Info().Msg("调用Heartbeat方法")
}

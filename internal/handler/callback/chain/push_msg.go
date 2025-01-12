package chain

import (
	"context"
	v1 "smallBot/api/gewe/v1"
	"smallBot/internal/sdk/gewe"

	"github.com/rs/zerolog/log"
)

type PushMsg struct {
	BaseHandler
	sdk *gewe.Gewe
}

func NewPushMsg(sdk *gewe.Gewe) *PushMsg {
	return &PushMsg{
		sdk: sdk,
	}
}

var _ Chain = (*PushMsg)(nil)

func (c *PushMsg) IsCanHandler(ctx context.Context, param v1.CollectRequest) bool {
	return true
}

func (c *PushMsg) HandlerRequest(ctx context.Context, param v1.CollectRequest) {
	log.Info().Msg("调用PushMsg->HandlerRequest方法")

	if c.IsCanHandler(ctx, param) {
		c.Process(ctx, param)
	}

	c.NextHandler.HandlerRequest(ctx, param)
}

func (c *PushMsg) Process(ctx context.Context, param v1.CollectRequest) {
	log.Info().Msg("调用PushMsg->Process方法")
}

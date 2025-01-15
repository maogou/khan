package chain

import (
	"context"
	v1 "smallBot/api/gewe/v1"
	"smallBot/internal/pkg/log"
	"smallBot/internal/sdk/gewe"
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
	log.C(ctx).Info().Msg("调用PushMsg->IsCanHandler方法")
	if len(c.sdk.Config().Sdk.Callback) > 0 {
		return true
	}

	return false
}

func (c *PushMsg) HandlerRequest(ctx context.Context, param v1.CollectRequest) {
	log.C(ctx).Info().Msg("调用PushMsg->HandlerRequest方法")

	if c.IsCanHandler(ctx, param) {
		_ = c.Process(ctx, param)
		if c.NextHandler != nil {
			c.NextHandler.HandlerRequest(ctx, param)
		}
	}

}

func (c *PushMsg) Process(ctx context.Context, param v1.CollectRequest) error {
	log.C(ctx).Info().Msg("调用PushMsg->Process方法")

	if err := c.sdk.PushMsg(ctx, param, c.sdk.Config().Sdk.Callback); err != nil {
		log.C(ctx).Error().Err(err).Msg("调用PushMsg方法失败")
		return err
	}
	log.C(ctx).Info().Msg("调用PushMsg方法成功")

	return nil
}

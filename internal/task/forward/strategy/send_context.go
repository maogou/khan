package strategy

import (
	"context"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/sdk/khan"
)

type SendStrategy interface {
	Send(ctx context.Context, req v1.SnsForwardRequest, timeline v1.TimelineObject) error
}

type SendContext struct {
	strategy SendStrategy
}

func (c *SendContext) Execute(ctx context.Context, req v1.SnsForwardRequest, timeline v1.TimelineObject) error {
	log.C(ctx).Info().Msg("开始执行发送策略")
	defer log.C(ctx).Info().Msg("发送策略执行完成")

	return c.strategy.Send(ctx, req, timeline)
}

func NewSendContext(contentStyle int, sdk *khan.Khan) *SendContext {
	var strategy SendStrategy
	switch contentStyle {
	case 1:
		strategy = &ImageStrategy{sdk: sdk}
	default:
		strategy = &DefaultStrategy{sdk: sdk}
	}
	return &SendContext{
		strategy: strategy,
	}
}

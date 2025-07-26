package chain

import (
	"context"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/sdk/khan"
	"maogou/khan/internal/task/forward/strategy"
)

type SendHandler struct {
	BaseHandler
	sdk *khan.Khan
}

func NewSendHandler(sdk *khan.Khan) Handler {
	return &SendHandler{
		sdk: sdk,
	}
}

var _ Handler = (*SendHandler)(nil)

func (s *SendHandler) Handle(ctx context.Context, pld *PipelineData) error {
	log.C(ctx).Info().Msg("开始执行消息转发.....")

	sendCtx := strategy.NewSendContext(pld.ContentStyle, s.sdk)
	if err := sendCtx.Execute(ctx, pld.Req, pld.Timeline); err != nil {
		log.C(ctx).Error().Err(err).Str("appid", pld.AppID).Msg("转发朋友圈发送失败")
		return err
	}

	log.C(ctx).Info().Str("appid", pld.AppID).Msg("转发朋友圈发送成功")

	pld.IsForwardOk = true

	if err := s.sdk.Rdb().Incr(ctx, constant.CurrentFavId+pld.AppID).Err(); err != nil {
		log.C(ctx).Error().Err(err).Msg("调用rdb.Incr方法,更新下一次转发的收藏favId失败,请手动更新redis中的值")
	}

	return nil
}

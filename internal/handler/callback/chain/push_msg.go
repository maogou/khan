package chain

import (
	"context"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/log"
	"smallBot/internal/sdk/khan"
)

type PushMsg struct {
	BaseHandler
	sdk *khan.Khan
}

func NewPushMsg(sdk *khan.Khan) *PushMsg {
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
	}

	c.ExecuteNext(ctx, param)

}

func (c *PushMsg) Process(ctx context.Context, param v1.CollectRequest) error {
	log.C(ctx).Info().Msg("调用PushMsg->Process方法")

	cKey := constant.WXLoginCache + param.AppId
	wxid, err := c.sdk.Rdb().Get(ctx, cKey).Result()
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("获取wxid失败")
		return err
	}

	if len(wxid) == 0 {
		log.C(ctx).Warn().Msg("获取wxid为空")
		return err
	}

	if err = c.sdk.PushMsg(ctx, param, c.sdk.Config().Sdk.Callback, wxid); err != nil {
		log.C(ctx).Error().Err(err).Msg("调用PushMsg方法失败")
		return err
	}
	log.C(ctx).Info().Msg("调用PushMsg方法成功")

	return nil
}

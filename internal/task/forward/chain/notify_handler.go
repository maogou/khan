package chain

import (
	"context"
	"maogou/khan/api/khan/v1/transform/message"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/sdk/khan"
	"strconv"
)

type NotifyHandler struct {
	BaseHandler
	sdk *khan.Khan
}

func NewNotifyHandler(sdk *khan.Khan) Handler {
	return &NotifyHandler{
		sdk: sdk,
	}
}

var _ Handler = (*NotifyHandler)(nil)

func (n *NotifyHandler) Handle(ctx context.Context, pld *PipelineData) error {
	log.C(ctx).Info().Str("appid", pld.AppID).Msg("开始执行转发朋友圈通知.....")

	nResp, err := n.sdk.PostText(
		ctx, message.PostTextRequest{
			AppId: pld.AppID,
			ToWxidList: []message.ToWxid{
				{
					Content: "收藏夹id(" + strconv.Itoa(pld.FavID) + ") 已经转发到你的朋友圈中发布内容，请及时查看!",
					ToWxid:  "filehelper",
					MsgType: 1,
				},
			},
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("发送通知消息失败")
		return err
	}

	if nResp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("nResp.Ret != 0 ->发送通知消息失败")
		return err
	}

	return nil
}

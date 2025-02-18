package chain

import (
	"context"
	"encoding/json"
	"encoding/xml"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/trade"
	"smallBot/internal/pkg/log"
	"smallBot/internal/sdk/khan"
	"strings"
	"time"
)

type RedPacket struct {
	BaseHandler
	sdk *khan.Khan
}

func NewRedPacket(sdk *khan.Khan) *RedPacket {
	return &RedPacket{
		sdk: sdk,
	}
}

var _ Chain = (*RedPacket)(nil)

func (r *RedPacket) HandlerRequest(ctx context.Context, param v1.CollectRequest) {
	log.C(ctx).Info().Msg("调用RedPacket->HandlerRequest方法")

	if r.IsCanHandler(ctx, param) {
		if err := r.Process(ctx, param); err != nil {
			log.C(ctx).Error().Err(err).Msg("调用RedPacket->Process方法失败")
		}
	}

	if r.NextHandler != nil {
		r.NextHandler.HandlerRequest(ctx, param)
	}
}

func (r *RedPacket) IsCanHandler(ctx context.Context, param v1.CollectRequest) bool {
	log.C(ctx).Info().Msg("调用RedPacket->IsCanHandler方法")
	var (
		rp  v1.WeChatMessageType49
		cxc v1.CallbackXmlContent
	)

	if param.MsgType != 49 {
		return false
	}

	if err := json.Unmarshal(param.Data, &cxc); err != nil {
		log.C(ctx).Error().Err(err).Msg("json解析49类型的消息失败")
		return false
	}

	xc := cxc.Content.String

	xc = strings.NewReplacer(
		"\t", "",
		"\n", "",
	).Replace(xc)

	if err := xml.Unmarshal([]byte(xc), &rp); err != nil {
		log.C(ctx).Error().Err(err).Msg("xml解析49类型的消息失败")
		return false
	}

	if rp.AppMessage.Type == "2001" {
		return true
	}

	return false
}

func (r *RedPacket) Process(ctx context.Context, param v1.CollectRequest) error {
	log.C(ctx).Info().Msg("调用RedPacket->Process方法")
	var (
		rp  v1.WeChatMessageType49
		cxc v1.CallbackXmlContent
	)

	if err := json.Unmarshal(param.Data, &cxc); err != nil {
		log.C(ctx).Error().Err(err).Msg("json解析49类型的消息失败")
		return err
	}

	xc := cxc.Content.String

	xc = strings.NewReplacer(
		"\t", "",
		"\n", "",
	).Replace(xc)

	if err := xml.Unmarshal([]byte(xc), &rp); err != nil {
		log.C(ctx).Error().Err(err).Msg("xml解析49类型的消息失败")
		return err
	}

	time.AfterFunc(
		50*time.Millisecond, func() {
			resp, err := r.sdk.OpenRedPacket(
				ctx, trade.RedPacketRequest{
					AppId:   param.AppId,
					GroupId: rp.FromUser,
					Xml:     xc,
				},
			)

			if err != nil {
				log.C(ctx).Error().Err(err).Msg("调用openRedPacket方法失败")
			} else {
				log.C(ctx).Info().Msgf("调用openRedPacket方法成功,返回值:%+v", resp)
			}
		},
	)

	return nil
}

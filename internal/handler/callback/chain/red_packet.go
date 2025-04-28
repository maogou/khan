package chain

import (
	"context"
	"encoding/json"
	"encoding/xml"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/trade"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/help"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/sdk/khan"
	"math/rand"
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

	r.ExecuteNext(ctx, param)
}

func (r *RedPacket) IsCanHandler(ctx context.Context, param v1.CollectRequest) bool {
	log.C(ctx).Info().Msg("调用RedPacket->IsCanHandler方法")
	var (
		rp  v1.WeChatMessageType49
		cxc v1.CallbackXmlContent
	)

	if err := json.Unmarshal(param.Data, &cxc); err != nil {
		log.C(ctx).Warn().Msg("json解析49类型的消息失败,是系统消息推送")
		return false
	}

	if cxc.MsgType != 49 {
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

	cKey := constant.WXPaymentCacheKey + param.AppId
	exist, err := r.sdk.Rdb().Get(ctx, cKey).Result()
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("获取自动抢红包配置失败")
		return false
	}

	shouldValue := help.Md5(constant.WXPaymentCacheKey + param.AppId)

	if exist != shouldValue {
		log.C(ctx).Warn().Str("shouldValue", shouldValue).Str("cacheValue", exist).Msg("自动收款未开启")
		return false
	}

	if rp.AppMessage.Type == constant.WXMsgTypeRedPacket {
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

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := time.Duration(500 + rng.Intn(501))

	time.AfterFunc(
		random*time.Millisecond, func() {
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

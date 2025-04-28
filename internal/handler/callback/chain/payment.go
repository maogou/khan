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

type Payment struct {
	BaseHandler
	sdk *khan.Khan
}

func (p *Payment) HandlerRequest(ctx context.Context, param v1.CollectRequest) {
	log.C(ctx).Info().Msg("调用Payment->HandlerRequest方法")

	if p.IsCanHandler(ctx, param) {
		if err := p.Process(ctx, param); err != nil {
			log.C(ctx).Error().Err(err).Msg("调用Payment->Process方法失败")
		}
	}

	p.ExecuteNext(ctx, param)
}

func (p *Payment) IsCanHandler(ctx context.Context, param v1.CollectRequest) bool {
	log.C(ctx).Info().Msg("调用Payment->IsCanHandler方法")

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
	exist, err := p.sdk.Rdb().Get(ctx, cKey).Result()
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("获取自动收款失败")
		return false
	}

	shouldValue := help.Md5(constant.WXPaymentCacheKey + param.AppId)

	if exist != shouldValue {
		log.C(ctx).Warn().Str("shouldValue", shouldValue).Str("cacheValue", exist).Msg("自动收款未开启")
		return false
	}

	if rp.AppMessage.Type == constant.WXMsgTypePayment {
		return true
	}

	return false
}

func (p *Payment) Process(ctx context.Context, param v1.CollectRequest) error {
	log.C(ctx).Info().Msg("调用Payment->Process方法")

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
			resp, err := p.sdk.CollectMoney(
				ctx, trade.CollectMoneyRequest{
					AppId: param.AppId,
					ToWid: rp.FromUser,
					Xml:   xc,
				},
			)

			if err != nil {
				log.C(ctx).Error().Err(err).Msg("调用CollectMoney方法失败")
			} else {
				log.C(ctx).Info().Msgf("调用CollectMoney方法成功,返回值:%+v", resp)
			}
		},
	)

	return nil
}

func NewPayment(sdk *khan.Khan) *Payment {
	return &Payment{
		sdk: sdk,
	}
}

var _ Chain = (*Payment)(nil)

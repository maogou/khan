package khan

import (
	"context"
	"maogou/khan/api/khan/v1/transform/trade"
	"maogou/khan/internal/pkg/log"
)

func (k *Khan) OpenRedPacket(ctx context.Context, req trade.RedPacketRequest) (*trade.RedPacketResponse, error) {
	resp, err := k.Client().R().SetBody(req).SetResult(&trade.RedPacketResponse{}).Post(redPacket)

	if err != nil {
		log.C(ctx).Error().Msg("调用openRedPacket方法失败")
		return nil, err
	}

	return resp.Result().(*trade.RedPacketResponse), nil
}

func (k *Khan) CollectMoney(ctx context.Context, req trade.CollectMoneyRequest) (*trade.CollectMoneyResponse, error) {
	resp, err := k.Client().R().SetBody(req).SetResult(&trade.CollectMoneyResponse{}).Post(collectMoney)

	if err != nil {
		log.C(ctx).Error().Msg("调用collectMoney方法失败")
		return nil, err
	}

	return resp.Result().(*trade.CollectMoneyResponse), nil
}

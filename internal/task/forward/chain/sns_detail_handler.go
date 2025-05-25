package chain

import (
	"context"
	"errors"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/sdk/khan"
)

type SnsDetailHandler struct {
	BaseHandler
	sdk *khan.Khan
}

func NewSnsDetailHandler(sdk *khan.Khan) Handler {
	return &SnsDetailHandler{
		sdk: sdk,
	}
}

var _ Handler = (*SnsDetailHandler)(nil)

func (sd *SnsDetailHandler) Handle(ctx context.Context, pld *PipelineData) error {
	log.C(ctx).Info().Str("appid", pld.AppID).Msg("开始执行获取sns详情.....")

	if len(pld.SnsID) == 0 {
		log.C(ctx).Warn().Str("appid", pld.AppID).Msg("snsID为空")
		return errors.New("snsID为空")
	}

	snsResp, err := sd.sdk.SnsDetail(
		ctx, v1.SnsDetailRequest{
			AppId: pld.AppID,
			SnsId: pld.SnsID,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsDetail方法失败")
		return err
	}

	if snsResp.Ret != 200 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用SnsDetail方法失败")
		return errors.New("ret != 0 ->调用SnsDetail方法失败")
	}

	pld.SnsDetail = &snsResp.Data

	return nil
}

package chain

import (
	"context"
	"errors"
	"smallBot/api/khan/v1/transform/sns"
	"smallBot/internal/pkg/log"
	"smallBot/internal/sdk/khan"
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
		ctx, sns.SnsDetailRequest{
			AppId:   pld.AppID,
			Decrypt: true,
			Id:      pld.SnsID,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsDetail方法失败")
		return err
	}

	if snsResp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用SnsDetail方法失败")
		return errors.New("ret != 0 ->调用SnsDetail方法失败")
	}

	pld.SnsDetail = snsResp

	return nil
}

package chain

import (
	"context"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/sdk/khan"
)

type FavDetailHandler struct {
	BaseHandler
	sdk *khan.Khan
}

func NewFavDetailHandler(sdk *khan.Khan) Handler {
	return &FavDetailHandler{
		sdk: sdk,
	}
}

var _ Handler = (*FavDetailHandler)(nil)

func (d *FavDetailHandler) Handle(ctx context.Context, pld *PipelineData) error {
	log.C(ctx).Info().Str("appid", pld.AppID).Msg("开始执行获取详情.....")

	detailResp, err := d.sdk.Detail(
		ctx, v1.FavorDetailRequest{
			AppId: pld.AppID,
			FavId: pld.FavID,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Str("appid", pld.AppID).Msg("调用fav Detail方法失败")
		return err
	}

	if detailResp.Ret != 200 {
		log.C(ctx).Error().Str("msg", detailResp.Msg).Str(
			"appid", pld.AppID,
		).Msg("调用fav FavDetail ->detailResp.Ret != 0 方法失败")
		return err
	}

	pld.FavDetail = &detailResp.Data

	return nil
}

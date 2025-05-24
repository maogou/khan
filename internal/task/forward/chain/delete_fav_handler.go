package chain

import (
	"context"
	"errors"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/sdk/khan"
)

type DeleteFavHandler struct {
	BaseHandler
	sdk *khan.Khan
}

func NewDeleteFavHandler(sdk *khan.Khan) Handler {
	return &DeleteFavHandler{
		sdk: sdk,
	}
}

var _ Handler = (*DeleteFavHandler)(nil)

func (df *DeleteFavHandler) Handle(ctx context.Context, pld *PipelineData) error {
	log.C(ctx).Info().Str("appid", pld.AppID).Msg("开始执行删除转发过的收藏.....")

	if !pld.IsForwardOk {
		log.C(ctx).Warn().Str("appid", pld.AppID).Msg("转发失败，不删除")

		return errors.New("转发失败，不删除")
	}

	delResp, err := df.sdk.Delete(
		ctx, v1.FavorDeleteRequest{
			AppId: pld.AppID,
			FavId: pld.FavID,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用Delete favor方法失败")
		return err
	}

	if delResp.Ret != 200 {
		log.C(ctx).Error().Str("msg", delResp.Msg).Msg("调用Delete favor方法失败")
		return errors.New("调用Delete favor方法失败")
	}

	return nil
}

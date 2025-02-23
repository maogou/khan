package chain

import (
	"context"
	"errors"
	"smallBot/internal/pkg/log"
	"smallBot/internal/sdk/khan"
)

type FavIdHandler struct {
	BaseHandler
	sdk *khan.Khan
}

func NewFavIdHandler(sdk *khan.Khan) Handler {
	return &FavIdHandler{
		sdk: sdk,
	}
}

var _ Handler = (*FavIdHandler)(nil)

func (f *FavIdHandler) Handle(ctx context.Context, pld *PipelineData) error {
	log.C(ctx).Info().Str("appid", pld.AppID).Msg("开始执行获取收藏夹ID.....")
	if len(pld.SyncResult.Data.List) == 0 {
		log.C(ctx).Warn().Msg("收藏夹为空")
		return errors.New("收藏夹为空")
	}

	for _, fav := range pld.SyncResult.Data.List {
		if fav.Flag == 1 && fav.FavId > 0 {
			continue
		}

		pld.FavID = fav.FavId

		break
	}

	if pld.FavID == 0 {
		log.C(ctx).Warn().Str("appid", pld.AppID).Msg("收藏夹中没有需要转发的朋友圈")
		return errors.New("收藏夹中没有需要转发的朋友圈")
	}

	return nil
}

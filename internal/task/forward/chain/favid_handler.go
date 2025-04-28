package chain

import (
	"context"
	"errors"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/sdk/khan"
	"strconv"
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

	if lLen, err := f.sdk.Rdb().LLen(ctx, constant.FavIds+pld.AppID).Result(); err != nil {
		log.C(ctx).Error().Err(err).Str("appid", pld.AppID).Msg("调用rdb.LLen方法失败")
		return err
	} else if lLen == 0 {
		log.C(ctx).Info().Str("appid", pld.AppID).Msg("队列中没有需要转发的朋友圈")
		return errors.New("队列中没有需要转发的朋友圈")
	}

	favId, err := f.sdk.Rdb().LPop(ctx, constant.FavIds+pld.AppID).Result()

	if err != nil {
		log.C(ctx).Error().Err(err).Str("appid", pld.AppID).Msg("调用rdb.LPop方法失败")
		return err
	}

	favInt, err := strconv.Atoi(favId)
	if err != nil {
		log.C(ctx).Error().Err(err).Str("appid", pld.AppID).Msg("调用strconv.Atoi方法失败")
		return err
	}

	pld.FavID = favInt

	if pld.FavID == 0 {
		log.C(ctx).Warn().Str("appid", pld.AppID).Msg("收藏夹中没有需要转发的朋友圈")
		return errors.New("收藏夹中没有需要转发的朋友圈")
	}

	return nil
}

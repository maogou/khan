package chain

import (
	"context"
	"errors"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/sdk/khan"
	"strconv"

	"github.com/spf13/cast"
)

const startFavId = 47876

type SyncHandler struct {
	BaseHandler
	sdk *khan.Khan
}

func NewSyncHandler(sdk *khan.Khan) Handler {
	return &SyncHandler{
		sdk: sdk,
	}
}

var _ Handler = (*SyncHandler)(nil)

func (s *SyncHandler) Handle(ctx context.Context, pld *PipelineData) error {
	log.C(ctx).Info().Str("appid", pld.AppID).Msg("开始执行同步收藏夹.....")

	var favIds = make([]string, 0)

	currentFavId, err := s.sdk.Rdb().Get(ctx, constant.CurrentFavId+pld.AppID).Result()
	if err != nil {
		log.C(ctx).Error().Err(err).Str("appid", pld.AppID).Msg("调用rdb.Get方法获取当前需要转发的收藏朋友圈id失败")
		return err
	}

	log.C(ctx).Info().Any("current_fav_id", currentFavId).Msg("当前需要转发的收藏id")

	if len(currentFavId) == 0 {
		log.C(ctx).Warn().Msg("缓存中当前需要转发的收藏id为空")
		return errors.New("缓存中当前需要转发的收藏id为空")
	}

	favId := cast.ToInt(currentFavId)
	if favId <= 0 {
		log.C(ctx).Warn().Msg("转换缓存中的收藏id后结果不大于0")
		return errors.New("转换缓存中的收藏id后结果不大于0")
	}

	detailResp, err := s.sdk.Detail(
		ctx, v1.FavorDetailRequest{
			AppId: pld.AppID,
			FavId: startFavId,
		},
	)

	log.C(ctx).Info().Any("detailResp", detailResp).Msg("获取收藏内容")

	if err != nil {
		log.C(ctx).Error().Err(err).Str("appid", pld.AppID).Int64("fav_id", startFavId).Msg("调用fav Detail方法失败")
		return err
	}

	if detailResp.Ret != 200 {
		log.C(ctx).Error().Str("appid", pld.AppID).Msg("调用fav FavDetail ->detailResp.Ret != 0 方法失败")
		return err
	}

	if detailResp.Data.Flag == 1 || len(detailResp.Data.Content) == 0 {
		log.C(ctx).Warn().Int64("fav_id", startFavId).Msg("当前收藏内容已经被删除")

		if rErr := s.sdk.Rdb().Incr(ctx, constant.CurrentFavId+pld.AppID).Err(); rErr != nil {
			log.C(ctx).Error().Err(rErr).Str(
				"appid", pld.AppID,
			).Msg("调用rdb.Incr方法,更新下一次转发的收藏favId失败,请手动更新redis中的值")
		}

		return errors.New("收藏内容已经被删除,等待进行下一次定时转发")
	}

	favIds = append(favIds, strconv.Itoa(detailResp.Data.FavId))

	if len(favIds) > 0 {
		if rErr := s.sdk.Rdb().RPush(ctx, constant.FavIds+pld.AppID, favIds).Err(); rErr != nil {
			log.C(ctx).Error().Err(rErr).Str("appid", pld.AppID).Msg("调用rdb.RPushX方法失败")
			return rErr
		}
	}

	return nil
}

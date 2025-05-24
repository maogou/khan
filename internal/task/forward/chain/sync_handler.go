package chain

import (
	"context"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/sdk/khan"
	"strconv"
	"time"
)

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

	var (
		syncKey string
		favIds  = make([]string, 0)
	)

	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	currentDayZeroUnix := midnight.Unix()

	for {
		lLen, err := s.sdk.Rdb().LLen(ctx, constant.FavIds+pld.AppID).Result()

		if err != nil {
			log.C(ctx).Error().Err(err).Str("appid", pld.AppID).Msg("调用rdb.LLen方法失败")
			return err
		}

		if lLen > 0 {
			break
		}

		syncResp, err := s.sdk.Sync(
			ctx, v1.FavorSyncRequest{
				AppId:   pld.AppID,
				SyncKey: syncKey,
			},
		)
		if err != nil {
			log.C(ctx).Error().Err(err).Str("appid", pld.AppID).Msg("调用favorSync方法失败")
			return err
		}

		if syncResp.Ret != 200 {
			log.C(ctx).Error().Str("msg", syncResp.Msg).Str(
				"appid", pld.AppID,
			).Msg("调用favorSync ->syncResp.Ret != 0 方法失败")
			return err
		}

		syncKey = syncResp.Data.SyncKey

		if len(syncResp.Data.List) == 0 {
			log.C(ctx).Info().Str("appid", pld.AppID).Msg("同步收藏夹完成")
			break
		}

		for _, item := range syncResp.Data.List {
			if item.Flag == 1 || currentDayZeroUnix-int64(item.UpdateTime) > 0 {
				continue
			}

			favIds = append(favIds, strconv.Itoa(item.FavId))
		}
	}

	if len(favIds) > 0 {
		if err := s.sdk.Rdb().RPush(ctx, constant.FavIds+pld.AppID, favIds).Err(); err != nil {
			log.C(ctx).Error().Err(err).Str("appid", pld.AppID).Msg("调用rdb.RPushX方法失败")
			return err
		}
	}

	return nil
}

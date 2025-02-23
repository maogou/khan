package chain

import (
	"context"
	"smallBot/api/khan/v1/transform/favor"
	"smallBot/internal/pkg/log"
	"smallBot/internal/sdk/khan"
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

	syncResp, err := s.sdk.Sync(
		ctx, favor.FavorSyncRequest{
			AppId:   pld.AppID,
			SyncKey: "",
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Str("appid", pld.AppID).Msg("调用favorSync方法失败")
		return err
	}

	if syncResp.Ret != 0 {
		log.C(ctx).Error().Str("appid", pld.AppID).Msg("调用favorSync ->syncResp.Ret != 0 方法失败")
		return err
	}

	pld.SyncResult = syncResp

	return nil
}

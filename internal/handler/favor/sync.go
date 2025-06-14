package favor

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/favor"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (f *FavorHandler) Sync(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用FavorHandler->Sync方法")

	var (
		req    v1.FavorSyncRequest
		result = v1.FavorSyncResponse{
			List: []v1.FavorSyncItem{},
		}
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := f.sdk.Sync(
		ctx, favor.FavorSyncRequest{
			AppId:   req.AppId,
			SyncKey: req.SyncKey,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用favorSync方法失败")
		response.Fail(ctx, errno.FavorSyncError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret !=0 ->调用favorSync方法失败")
		response.Fail(ctx, errno.FavorSyncError)
		return
	}

	result.SyncKey = resp.Data.KeyBuf.Buffer

	for _, v := range resp.Data.List {
		result.List = append(
			result.List, v1.FavorSyncItem{
				FavId:      v.FavId,
				Type:       v.Type,
				Flag:       v.Flag,
				UpdateTime: v.UpdateTime,
			},
		)
	}

	response.Success(ctx, result)
}

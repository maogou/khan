package favor

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (f *FavorHandler) Sync(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用FavorHandler->Sync方法")

	var (
		req v1.FavorSyncRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := f.sdk.Sync(ctx, req)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用favorSync方法失败")
		response.Fail(ctx, errno.FavorSyncError)
		return
	}

	if resp.Ret != 200 {
		log.C(ctx).Error().Str("msg", resp.Msg).Msg("ret !=0 ->调用favorSync方法失败")
		response.Fail(ctx, errno.FavorSyncError)
		return
	}

	response.Success(ctx, resp.Data)
}

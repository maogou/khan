package favor

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (f *FavorHandler) Delete(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用FavorHandler->Delete方法")

	var req v1.FavorDeleteRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := f.sdk.Delete(ctx, req)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用Delete方法失败")
		response.Fail(ctx, errno.FavorDeleteError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Str("msg", resp.Msg).Msg("ret !=0 ->调用Delete方法失败")
		response.Fail(ctx, errno.FavorDeleteError)
		return
	}

	response.SuccessMsg(ctx, "删除成功")
}

package favor

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/favor"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

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

	resp, err := f.sdk.Delete(
		ctx, favor.FavorDeleteRequest{
			AppId: req.AppId,
			FavId: req.FavId,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用Delete方法失败")
		response.Fail(ctx, errno.FavorDeleteError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret !=0 ->调用Delete方法失败")
		response.Fail(ctx, errno.FavorDeleteError)
		return
	}

	response.SuccessMsg(ctx, "删除成功")
}

package favor

import (
	"github.com/gin-gonic/gin"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/label"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
)

func (f *FavorHandler) Detail(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用FavorHandler->Detail方法")

	var (
		req    v1.FavorDetailRequest
		result v1.FavorDetailResponse
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := f.sdk.Detail(
		ctx, label.FavorDetailRequest{
			AppId: req.AppId,
			FavId: req.FavId,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用Detail方法失败")
		response.Fail(ctx, errno.FavorDetailError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret !=0 ->调用Detail方法失败")
		response.Fail(ctx, errno.FavorDetailError)
		return
	}

	if len(resp.Data.ObjectList) == 0 {
		response.Success(ctx, result)
		return
	}

	response.Success(
		ctx, v1.FavorDetailResponse{
			FavId:      resp.Data.ObjectList[0].FavId,
			Status:     resp.Data.ObjectList[0].Status,
			Flag:       resp.Data.ObjectList[0].Flag,
			UpdateTime: resp.Data.ObjectList[0].UpdateTime,
			Content:    resp.Data.ObjectList[0].Object,
		},
	)
}

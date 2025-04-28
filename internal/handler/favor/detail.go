package favor

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/favor"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
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
		ctx, favor.FavorDetailRequest{
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

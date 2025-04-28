package sns

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/sns"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (s *SnsHandler) Like(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->Like方法")

	var req v1.SnsLikeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := s.sdk.SnsLike(
		ctx, sns.SnsLikeRequest{
			AppId:  req.AppId,
			Type:   1,
			Id:     req.SnsId,
			ToWxid: req.Wxid,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsLike方法失败")
		response.Fail(ctx, errno.SnsLikeError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret !=0 ->调用SnsLike方法失败")
		response.Fail(ctx, errno.SnsLikeError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("BaseResponse.Ret !=0 ->调用SnsLike方法失败")
		response.Fail(ctx, errno.SnsLikeError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")
}

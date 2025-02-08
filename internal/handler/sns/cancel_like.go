package sns

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/sns"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (s *SnsHandler) CancelLike(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->CancelLike方法")

	var req v1.SnsCancelLikeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := s.sdk.SnsCancelLike(
		ctx, sns.SnsCancelLikeRequest{
			AppId: req.AppId,
			Type:  5,
			Id:    req.SnsId,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsCancelLike方法失败")
		response.Fail(ctx, errno.SnsCancelLikeError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret !=0 ->调用SnsCancelLike方法失败")
		response.Fail(ctx, errno.SnsCancelLikeError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")
}

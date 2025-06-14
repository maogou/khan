package sns

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/sns"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (s *SnsHandler) DeleteComment(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->DeleteComment方法")

	var req v1.SnsDelCommentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := s.sdk.SnsDeleteComment(
		ctx, sns.SnsDeleteCommentRequest{
			AppId:     req.AppId,
			Id:        req.SnsId,
			Type:      4,
			CommnetId: req.CommentId,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsDeleteComment方法失败")
		response.Fail(ctx, errno.SnsDeleteCommentError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret !=0 ->调用SnsDeleteComment方法失败")
		response.Fail(ctx, errno.SnsDeleteCommentError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("BaseResponse.Ret !=0 ->调用SnsDeleteComment方法失败")
		response.Fail(ctx, errno.SnsDeleteCommentError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")
}

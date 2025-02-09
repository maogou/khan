package sns

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/sns"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (s *SnsHandler) Delete(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->Delete方法")

	var req v1.SnsDeleteRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := s.sdk.SnsDelete(
		ctx, sns.SnsDeleteRequest{
			AppId: req.AppId,
			Id:    req.SnsId,
			Type:  1,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsDelete方法失败")
		response.Fail(ctx, errno.SnsDeleteError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("ret != 0 ->调用SnsDelete方法失败")
		response.Fail(ctx, errno.SnsDeleteError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Any(
			"msg_err", resp.Data.BaseResponse.ErrMsg,
		).Msg("BaseResponse.Ret !=0 ->调用SnsDelete方法失败")
		response.Fail(ctx, errno.SnsDeleteError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")
}

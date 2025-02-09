package sns

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/sns"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/samber/lo"

	"github.com/gin-gonic/gin"
)

func (s *SnsHandler) SetPrivacy(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->SetPrivacy方法")

	var req v1.SnsSetPrivacyRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := s.sdk.SnsSetPrivacy(
		ctx, sns.SnsSetPrivacyRequest{
			AppId: req.AppId,
			Id:    req.SnsId,
			Type:  lo.Ternary(req.Open, 3, 2),
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsHandler->SetPrivacy方法失败")
		response.Fail(ctx, errno.SnsSetPrivacyError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用SnsHandler->SetPrivacy方法失败")
		response.Fail(ctx, errno.SnsSetPrivacyError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")
}

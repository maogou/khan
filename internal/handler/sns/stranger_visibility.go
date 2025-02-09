package sns

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/sns"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func (s *SnsHandler) StrangerVisibility(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->StrangerVisibility方法")

	var req v1.SnsStrangerVisibilityEnabledRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	id := lo.Ternary(req.Enabled, 6272, 6273)

	resp, err := s.sdk.SnsStrangerVisibilityEnabled(
		ctx, sns.SnsStrangerVisibilityEnabledRequest{
			AppId: req.AppId,
			Id:    id,
			Value: 72,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsSetPrivacyScope方法失败")
		response.Fail(ctx, errno.SnsStrangerVisibilityEnabledError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用SnsSetPrivacyScope方法失败")
		response.Fail(ctx, errno.SnsStrangerVisibilityEnabledError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")

}

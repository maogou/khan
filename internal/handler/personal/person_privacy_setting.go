package personal

import (
	v1 "smallBot/api/gewe/v1"
	"smallBot/api/gewe/v1/transform"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/samber/lo"

	"github.com/gin-gonic/gin"
)

func (p *PersonalHandler) PrivacySetting(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用PersonalHandler->PrivacySetting方法")

	var req v1.PersonalPrivacySettingRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := p.sdk.PersonalPrivacySetting(
		ctx, transform.PersonalPrivacySettingRequest{
			Appid: req.AppId,
			Id:    req.Option,
			Value: lo.Ternary(req.Open, 1, 0),
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用sdk->PersonalPrivacySetting方法失败")
		response.Fail(ctx, errno.PersonalPrivacySettingError)
		return
	}

	response.SuccessMsg(ctx, lo.Ternary(resp.Ret == 0, "设置成功", "设置失败"))
}

package login

import (
	v1 "smallBot/api/gewe/v1"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (l *LoginHandler) SetCallback(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用LoginHandler->SetCallback方法")

	var req v1.SetCallbackRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	_, err := l.sdk.Client().R().Post(req.CallbackUrl)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SetCallback方法失败")
		response.Fail(ctx, errno.SetCallbackError)
		return
	}

	log.C(ctx).Info().Str("old_callback_url", l.sdk.Config().Sdk.Callback).Msg("原来的callback_url")

	l.sdk.Config().Sdk.Callback = req.CallbackUrl

	log.C(ctx).Info().Str("new_callback_url", l.sdk.Config().Sdk.Callback).Msg("设置后的callback_url")

	response.SuccessMsg(ctx, "设置成功")
}

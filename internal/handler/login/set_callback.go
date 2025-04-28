package login

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

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

	cKey := constant.WXCallbackCache + req.AppId

	if err = l.sdk.Rdb().Set(ctx, cKey, req.CallbackUrl, 0).Err(); err != nil {
		log.C(ctx).Error().Err(err).Msg("设置回调保存redis失败")
		response.Fail(ctx, errno.SetTokenCallbackError)
	}

	response.SuccessMsg(ctx, "设置成功")
}

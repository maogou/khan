package login

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/login"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (l *LoginHandler) Logout(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用LoginHandler->Logout方法")

	var req v1.LogoutRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := l.sdk.Logout(
		ctx, login.LogoutRequest{
			Appid: req.AppId,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用Logout方法失败")
		response.Fail(ctx, errno.LogoutError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret !=0 ->调用Logout方法失败")
		response.Fail(ctx, errno.LogoutError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("BaseResponse.Ret !=0 ->调用Logout方法失败")
		response.Fail(ctx, errno.LogoutError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")
}

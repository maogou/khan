package login

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (l *LoginHandler) CreateApp(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用LoginHandler->CreateApp方法")

	caResp, err := l.sdk.CreateApp(
		ctx, v1.CreateAppRequest{
			Country:    constant.Country,
			DeviceName: constant.DeviceName,
			Model:      constant.Model,
			SdkVer:     constant.ProtoVersion,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("创建appId,请求CreateApp失败")
		response.Fail(ctx, errno.CreateAppErr)
		return
	}

	if caResp.Ret != 0 {
		log.C(ctx).Warn().Str("err_msg", caResp.MsgErr).Msg("调用创建appId失败")
		response.Fail(ctx, errno.CreateAppErr)
		return
	}

	response.Success(
		ctx,
		v1.LoginCreateApp{
			AppId: caResp.Data.AppId,
			Tip:   "appid是唯一标识,请妥善保管,请勿外泄!",
		},
	)

}

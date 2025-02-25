package login

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (l *LoginHandler) Reconnection(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用LoginHandler->Reconnection方法")

	var req v1.AppIdRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	sResp, err := l.sdk.SecAutoAuth(
		ctx, v1.SecAutoAuthRequest{
			AppId:  req.AppId,
			SdkVer: constant.ProtoVersion,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SecAutoAuth接口失败")
		response.Fail(ctx, errno.SecAutoAuthError)
		return
	}

	if sResp.Ret != 0 {
		log.C(ctx).Error().Msg("SecAutoAuth接口返回ret != 0")
		response.Fail(ctx, errno.SecAutoAuthError)
		return
	}

	if sResp.CommonResponse.Ret != 0 {
		log.C(ctx).Error().Msg("SecAutoAuth接口返回CommonResponse.Ret != 0")
		response.Fail(ctx, errno.SecAutoAuthError)
		return
	}

	port := strconv.Itoa(l.sdk.Config().Port)

	lResp, err := l.sdk.LongOpen(
		ctx, v1.LongOpenRequest{
			AppId:      req.AppId,
			CleanCache: true,
			Host:       "http://127.0.0.1:" + port + "/api/v1/collect",
			Timeout:    60,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用开启长连接失败")
		response.Fail(ctx, errno.LongOpenError)
		return
	}

	if lResp.Ret != 0 && lResp.MsgErr != constant.WXLongAlreadyConnect {
		log.C(ctx).Error().Msg("开启长连接接口返回ret != 0")
		response.Fail(ctx, errno.LongOpenError)
		return
	}

	response.SuccessMsg(ctx, "ok")
}

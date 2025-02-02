package login

import (
	v1 "smallBot/api/gewe/v1"
	"smallBot/api/gewe/v1/transform"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (l *LoginHandler) LoginQrCode(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用LoginHandler->LoginQrCode方法")

	var req v1.GetLoginQrCodeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := l.sdk.LoginQrCode(
		ctx, transform.GetLoginQrCodeRequest{
			AppId: req.AppId,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用LoginQrCode方法失败")
		response.Fail(ctx, errno.LoginQrCodeError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用LoginQrCode方法失败")
		response.Fail(ctx, errno.LoginQrCodeError)
		return
	}

	response.Success(
		ctx, v1.GetLoginQrCodeResponse{
			AppId:       req.AppId,
			QrData:      resp.Data.Url,
			QrImgBase64: resp.Data.Base64,
			Uuid:        resp.Data.Uuid,
			NewKey:      resp.Data.Nkey,
		},
	)
}

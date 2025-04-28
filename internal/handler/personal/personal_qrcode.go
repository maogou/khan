package personal

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/personal"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (p *PersonalHandler) Qrcode(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用PersonalHandler->Qrcode方法")

	var req v1.PersonalQrcodeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := p.sdk.PersonalQrcode(
		ctx, personal.PersonalQrcodeRequest{
			AppId:   req.AppId,
			ProxyIp: req.ProxyIp,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用PersonalQrcode方法失败")
		response.Fail(ctx, errno.PersonalQrcodeError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用PersonalQrcode方法失败")
		response.Fail(ctx, errno.PersonalQrcodeError)
		return
	}

	response.Success(
		ctx, v1.PersonalQrcodeResponse{
			QrCode: resp.Data.Qrcode.Buffer,
		},
	)
}

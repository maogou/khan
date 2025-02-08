package personal

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/personal"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/samber/lo"

	"github.com/gin-gonic/gin"
)

func (p *PersonalHandler) UploadHdImg(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用PersonalHandler->UploadHdImg方法")

	var req v1.PersonalUploadHeadImgRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := p.sdk.PersonalUploadHeadImg(
		ctx, personal.PersonalUploadHdHeadImgRequest{
			Appid:    req.AppId,
			FileLink: req.HeadImgUrl,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用PersonalUploadHeadImg方法失败")
		response.Fail(ctx, errno.PersonalUploadHeadImgError)
		return
	}

	response.SuccessMsg(ctx, lo.Ternary(resp.Ret == 0, "更新头像成功", "更新头像失败"))
}

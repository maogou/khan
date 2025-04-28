package personal

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/personal"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (p *PersonalHandler) UpdateProfile(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用PersonalHandler->UpdateProfile方法")

	var req v1.PersonalUpdateProfileRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := p.sdk.PersonalUpdateProfile(
		ctx, personal.PersonalUpdateProfileRequest{
			Appid:    req.AppId,
			City:     req.City,
			Country:  req.Country,
			NickName: req.NickName,
			Province: req.Province,
			Sex:      *req.Sex,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用sdk->PersonalUpdateProfile方法失败")
		response.Fail(ctx, errno.PersonalUpdateProfileError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("调用sdk->PersonalUpdateProfile方法resp.Ret != 0失败")
		response.Fail(ctx, errno.PersonalUpdateProfileError)
		return
	}

	response.SuccessMsg(ctx, "更新个人资料成功")
}

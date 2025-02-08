package personal

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/personal"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (p *PersonalHandler) Profile(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用PersonalHandler->Profile方法")

	var req v1.PersonalProfileRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := p.sdk.PersonalProfile(
		ctx, personal.PersonalProfileRequest{
			AppId:   req.AppId,
			ProxyIp: req.ProxyIp,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用PersonalProfile方法失败")
		response.Fail(ctx, errno.PersonalProfileError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用PersonalProfile方法失败")
		response.Fail(ctx, errno.PersonalProfileError)
		return
	}

	response.Success(
		ctx, v1.PersonalProfileResponse{
			Alias:           resp.Data.UserInfo.Alias,
			Wxid:            resp.Data.UserInfo.UserName.String,
			NickName:        resp.Data.UserInfo.NickName.String,
			Mobile:          "",
			Uin:             resp.Data.UserInfo.BindUin,
			Sex:             resp.Data.UserInfo.Sex,
			Province:        resp.Data.UserInfo.Province,
			City:            resp.Data.UserInfo.Province,
			Signature:       resp.Data.UserInfo.Signature,
			Country:         resp.Data.UserInfo.Country,
			BigHeadImgUrl:   resp.Data.UserInfoExt.BigHeadImgUrl,
			SmallHeadImgUrl: resp.Data.UserInfoExt.SmallHeadImgUrl,
			SnsBgImg:        "",
		},
	)

	return
}

package login

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/login"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (l *LoginHandler) CheckLogin(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用LoginHandler->CheckLogin方法")
	var req v1.CheckLoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := l.sdk.CheckLoginQrCode(
		ctx, login.CheckLoginRequest{
			Appid: req.AppId,
			Nkey:  req.Newkey,
			Uuid:  req.Uuid,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用CheckLoginQrCode方法失败")
		response.Fail(ctx, errno.CheckLoginQrCodeError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用CheckLoginQrCode方法失败")
		response.Fail(ctx, errno.CheckLoginQrCodeError)
		return
	}

	wxid := resp.Data.LoginInfo.AcctSectResp.UserName
	lcKey := constant.WXLoginCache + req.AppId
	if err = l.sdk.Rdb().Set(ctx, lcKey, wxid, 0).Err(); err != nil {
		log.C(ctx).Error().Err(err).Msg("设置wx登录缓存失败")
		response.Fail(ctx, errno.SetWxLoginCacheError)
		return
	}

	response.Success(
		ctx, v1.CheckLoginResponse{
			Uuid:        resp.Data.StatusInfo.Uuid,
			HeadImgUrl:  resp.Data.StatusInfo.HeadImgUrl,
			NickName:    resp.Data.StatusInfo.NickName,
			ExpiredTime: resp.Data.StatusInfo.ExpiredTime,
			Status:      resp.Data.StatusInfo.Status,
			LoginInfo: v1.CheckLoginLoginInfo{
				Uin:      resp.Data.LoginInfo.AuthSectResp.Uin,
				Wxid:     wxid,
				NickName: resp.Data.StatusInfo.NickName,
				Mobile:   "",
				Alias:    resp.Data.LoginInfo.AcctSectResp.Alias,
			},
		},
	)
}

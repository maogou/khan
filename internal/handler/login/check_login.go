package login

import (
	"context"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/login"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/help"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"
	"maogou/khan/internal/sdk/khan"

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

	lic, err := help.License(ctx, l.sdk.Rdb())
	cCtx := ctx.Copy()
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("获取license失败")
		go verifyFailLogout(cCtx, l.sdk, req.AppId)
		response.Fail(ctx, errno.GetLicenseError)
		return
	}

	alias := resp.Data.LoginInfo.AcctSectResp.Alias
	log.C(ctx).Info().Str("wxid", wxid).Str("alias", alias).Msg("获取到的微信号")

	if lic.Cus != wxid && lic.Cus != resp.Data.LoginInfo.AcctSectResp.Alias {
		log.C(ctx).Error().Err(err).Msg("登录的微信号的和授权的微信号不相符")
		go verifyFailLogout(cCtx, l.sdk, req.AppId)
		response.Fail(ctx, errno.AuthWixLicenseError)
		return
	}

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

func verifyFailLogout(ctx context.Context, sdk *khan.Khan, appId string) {
	log.C(ctx).Error().Msg("verifyFailLogout==>授权微信号和登录微信号不相符")
	resp, err := sdk.Logout(
		ctx, login.LogoutRequest{
			Appid: appId,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("verifyFailLogout=>调用Logout方法失败")
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("verifyFailLogout=>ret !=0 ->调用Logout方法失败")
		return
	}
}

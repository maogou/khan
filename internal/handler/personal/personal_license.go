package personal

import (
	"encoding/json"
	"math"
	"slices"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/help"
	"smallBot/internal/pkg/license"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (p *PersonalHandler) License(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用PersonalHandler->License方法")

	var (
		pKey       string
		permission v1.Permission
		pln        = make([]v1.PersonalLicensePermission, 0)
	)

	token := ctx.GetHeader(constant.KhanToken)

	decrypted, err := help.AesDecrypt(token, []byte(constant.AesWXidKey))

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("token 解密失败")
		response.Fail(ctx, errno.DecryptTokenFail)
		return
	}

	appid := decrypted[len(constant.AppName):len(decrypted)]

	keys := []string{constant.License, constant.LicenseKey}
	vals, err := p.sdk.Rdb().MGet(ctx, keys...).Result()

	lb, ok := vals[0].(string)
	if !ok || len(lb) == 0 {
		log.C(ctx).Error().Str("url", ctx.Request.URL.Path).Str(
			"client_ip", ctx.ClientIP(),
		).Msg("Redis 获取授权许可证为空")
		response.SuccessMsg(ctx, "Redis 获取授权许可证为空")
		return
	}

	pKey, ok = vals[1].(string)
	if !ok || len(pKey) == 0 {
		log.C(ctx).Error().Str("url", ctx.Request.URL.Path).Str(
			"client_ip", ctx.ClientIP(),
		).Msg("Redis 获取授权许可证key为空")
		response.SuccessMsg(ctx, "Redis 获取授权许可证key为空")
		return
	}

	originPKey := pKey

	pKey = strings.NewReplacer(
		constant.License37, "+",
		constant.License73, "/",
		constant.License919, "=",
	).Replace(pKey)

	lic, err := license.Parse(pKey, lb)

	if err = json.Unmarshal(lic.Dat, &permission); err != nil {
		log.C(ctx).Error().Err(err).Msg("json解析授权许可证失败")
		response.SuccessMsg(ctx, "json解析授权许可证失败")
		return
	}
	log.C(ctx).Info().Any("p", permission.AppId).Any("appid", appid).Msg("sss")
	if !slices.Contains(permission.AppId, appid) {
		log.C(ctx).Warn().Str("appid", appid).Msg("请求的携带的token对应的wxid未授权，无法使用khan服务")
		response.SuccessMsg(ctx, "请求的携带的token对应的wxid未授权，无法使用khan服务")
		return
	}

	for key, v := range permission.Permission {
		pln = append(
			pln, v1.PersonalLicensePermission{
				Name:  key,
				Value: v,
			},
		)
	}

	response.Success(
		ctx, v1.PersonalLicenseResponse{
			AppId:          permission.AppId,
			Token:          permission.Token,
			CreateTime:     time.Unix(lic.Iat.Unix(), 0).Format(time.DateTime),
			ExpireTime:     time.Unix(lic.Exp.Unix(), 0).Format(time.DateTime),
			RemainDays:     int(math.Floor(float64(lic.Exp.Unix()-lic.Iat.Unix()) / float64(86400))),
			LicenseKey:     originPKey,
			LicenseContent: lb,
			Permission:     pln,
			LicenseType:    lic.Typ,
		},
	)
}

package middleware

import (
	"encoding/json"
	"slices"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/config"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/help"
	"smallBot/internal/pkg/license"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var excludePaths = []string{
	"/",
	"/api/v1/collect",
	"/v1/api/download",
	"/api/v1/callback",
	"/v1/api/personal/license",
	"/v1/api/login/createApp",
	"/v1/api/setting/license",
	"/v1/api/monitor",
}

func VerifyLicense(rdb *redis.Client, conf config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var (
			p    v1.Permission
			pKey string
		)

		url := ctx.Request.URL.Path
		paths := strings.Split(url, "/")

		if slices.Contains(excludePaths, url) || strings.Contains(
			conf.Sdk.Callback, url,
		) || strings.Contains(strings.Trim(url, "/"), constant.DebugPpro) {
			ctx.Next()
			return
		}

		pKey, err := rdb.Get(ctx, constant.LicenseKey).Result()
		if err != nil {
			log.C(ctx).Error().Err(err).Msg("Redis 获取授权许可证key为空")
			response.SuccessMsg(ctx, "Redis 获取授权许可证key为空")
			ctx.Abort()
			return
		}

		lb, err := rdb.Get(ctx, constant.License).Result()

		if err != nil {
			log.C(ctx).Error().Err(err).Msg("Redis 获取授权许可证为空")
			response.SuccessMsg(ctx, "Redis 获取授权许可证为空")
			ctx.Abort()
			return
		}

		pKey = strings.NewReplacer(
			constant.License37, "+",
			constant.License73, "/",
			constant.License919, "=",
		).Replace(pKey)

		lic, err := license.Parse(pKey, lb)

		if err = json.Unmarshal(lic.Dat, &p); err != nil {
			log.C(ctx).Error().Err(err).Msg("json解析授权许可证失败")
			ctx.Abort()
			response.SuccessMsg(ctx, "json解析授权许可证失败")
			return
		}

		if lic.Expired() {
			log.C(ctx).Error().Msg("许可证已过期")
			ctx.Abort()
			response.SuccessMsg(ctx, "许可证已过期")
			return
		}

		log.C(ctx).Info().Any("paths", paths).Any("permission", p).Msg("paths and permission")
		if len(paths) > 3 && !strings.Contains(conf.Sdk.Callback, url) {
			if value, pOk := p.Permission[paths[3]]; pOk && value != 1 {
				log.C(ctx).Info().Str("paths[3]", paths[3]).Any("permission", p).Msg("无此接口的访问权限")
				ctx.Abort()
				response.SuccessMsg(ctx, "无此接口的访问权限,请联系软件作者!")
				return
			}
		}

		token := ctx.GetHeader(constant.KhanToken)
		if len(token) == 0 {
			log.C(ctx).Error().Msg("请求头中未携带token")
			ctx.Abort()

			response.SuccessMsg(ctx, "请求头中未携带token")
			return
		}

		decrypted, err := help.AesDecrypt(token, []byte(constant.AesWXidKey))
		if err != nil {
			log.C(ctx).Error().Err(err).Msg("token 解密失败")
			ctx.Abort()
			response.SuccessMsg(ctx, "token 解密失败")
			return
		}

		if len(p.AppId) == 1 && p.AppId[0] == constant.KhanFree {
			ctx.Next()
			return
		}

		appid := strings.TrimLeft(decrypted, constant.AppName)

		if !slices.Contains(p.AppId, appid) {
			log.C(ctx).Warn().Msg("请求的携带的token对应的wxid未授权，无法使用khan服务")
			ctx.Abort()
			response.SuccessMsg(ctx, "请求的携带的token对应的wxid未授权，无法使用khan服务")
			return
		}

		ctx.Next()

	}
}

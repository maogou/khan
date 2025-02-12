package middleware

import (
	"encoding/json"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/license"
	"smallBot/internal/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type PostRequest struct {
	AppID string `json:"appid"`
	AppId string `json:"appId"`
}

func VerifyLicense(rdb *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var (
			p    v1.Permission
			pKey string
		)

		keys := []string{constant.License, constant.LicenseKey}
		vals, err := rdb.MGet(ctx, keys...).Result()

		lb, ok := vals[0].(string)
		if !ok || len(lb) == 0 {
			log.Ctx(ctx).Error().Str("url", ctx.Request.URL.Path).Str(
				"client_ip", ctx.ClientIP(),
			).Msg("Redis 获取授权许可证为空")
			response.SuccessMsg(ctx, "Redis 获取授权许可证为空")
			ctx.Abort()
			return
		}

		pKey, ok = vals[1].(string)
		if !ok || len(pKey) == 0 {
			log.Ctx(ctx).Error().Str("url", ctx.Request.URL.Path).Str(
				"client_ip", ctx.ClientIP(),
			).Msg("Redis 获取授权许可证key为空")
			response.SuccessMsg(ctx, "Redis 获取授权许可证key为空")
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
			log.Error().Err(err).Msg("json解析授权许可证失败")
			ctx.Abort()
			response.SuccessMsg(ctx, "json解析授权许可证失败")
			return
		}

		if lic.Expired() {
			log.Error().Msg("许可证已过期")
			ctx.Abort()
			response.SuccessMsg(ctx, "许可证已过期")
			return
		}

		url := ctx.Request.URL.Path
		paths := strings.Split(url, "/")

		if len(paths) > 3 {
			if value, pOk := p.Permission[paths[3]]; pOk && value != 1 {
				log.Ctx(ctx).Info().Str("url", url).Any("permission", p).Msg("无此接口的访问权限")
				ctx.Abort()
				response.SuccessMsg(ctx, "无此接口的访问权限,请联系软件作者(wechat: wxksky2022)")
				return
			}

			//获取post json请求中的appid
			var req PostRequest
			if err := ctx.ShouldBindJSON(&req); err != nil {
				log.Ctx(ctx).Error().Err(err).Msg("解析 JSON 请求体失败")
				ctx.Abort()
				response.SuccessMsg(ctx, "解析 JSON 请求体失败")
				return
			}

			appid := req.AppId
			log.Ctx(ctx).Info().Str("appid=========>", appid).Msg("ssssss")

		}

		ctx.Next()
	}
}

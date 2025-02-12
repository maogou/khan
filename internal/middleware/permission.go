package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/license"
	"smallBot/internal/pkg/response"
)

func Permission(rdb *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			l license.License
			p = make(map[string]int)
		)
		lb, err := rdb.Get(ctx, constant.License).Result()
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("redis获取授权许可证失败")
			ctx.Abort()
			response.SuccessMsg(ctx, "redis获取授权许可证失败")
			return
		}

		if err = json.Unmarshal([]byte(lb), &l); err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("json解析授权许可证失败")
			ctx.Abort()
			response.SuccessMsg(ctx, "json解析授权许可证失败")
			return
		}

		if l.Expired() {
			log.Error().Msg("许可证已过期")
			ctx.Abort()
			response.SuccessMsg(ctx, "许可证已过期")
			return
		}

		if err = json.Unmarshal(l.Dat, &p); err != nil {
			log.Error().Err(err).Msg("json解析授权许可证失败")
			ctx.Abort()
			response.SuccessMsg(ctx, "json解析授权许可证失败")
			return
		}

		//if nLic.Expired() {
		//	log.Error().Msg("许可证已过期")
		//	c.Abort()
		//	response.SuccessMsg(c, "许可证已过期")
		//	return
		//}

		ctx.Next()
	}
}

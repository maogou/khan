package middleware

import (
	"smallBot/internal/config"
	"smallBot/internal/pkg/license"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func VerifyLicense(conf config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		nLic, err := license.Parse(conf.Sdk.Key, conf.Sdk.License)
		if err != nil {
			log.Error().Err(err).Msg("许可证校验失败")
			c.Abort()
			response.SuccessMsg(c, "许可证校验失败")
			return
		}

		if nLic.Expired() {
			log.Error().Msg("许可证已过期")
			c.Abort()
			response.SuccessMsg(c, "许可证已过期")
			return
		}

		c.Next()
	}
}

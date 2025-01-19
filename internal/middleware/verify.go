package middleware

import (
	"os"
	"path/filepath"
	"smallBot/internal/config"
	"smallBot/internal/pkg/license"
	"smallBot/internal/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func VerifyLicense(conf config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := os.Stat(conf.Sdk.License); err != nil {
			log.Error().Err(err).Msg("许可证不存在")
			c.Abort()
			response.SuccessMsg(c, "许可证不存在")
			return
		}

		pKey := filepath.Base(conf.Sdk.License)

		pKey = strings.ReplaceAll(pKey, "37", "+")
		pKey = strings.ReplaceAll(pKey, "73", "/")

		nLic, err := license.Parse(pKey, conf.Sdk.License)
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

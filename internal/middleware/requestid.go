package middleware

import (
	"smallBot/internal/constant"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := xid.New().String()
		// 创建带有请求日志 ID 的 logger
		ctxLogger := log.With().Str(constant.QID, requestId).Logger()
		//把带有qid的zerolog 放入到上下文
		c.Set(constant.ZEROLOG, ctxLogger)
		c.Set(constant.QID, requestId)
		c.Writer.Header().Set(constant.XID, requestId)

		c.Next()
	}
}

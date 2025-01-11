package middleware

import (
	"smallBot/internal/constant"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := xid.New().String()
		c.Set(constant.QID, requestId)

		c.Next()
	}
}

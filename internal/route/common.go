package router

import (
	"smallBot/internal/constant"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
)

// initCommonRoute 初始化404 & 405对应的路由
func initCommonRoute(engine *gin.Engine) {
	engine.NoRoute(
		func(ctx *gin.Context) {
			response.Fail(ctx, errno.PageNotFoundError)
		},
	)

	engine.NoMethod(
		func(ctx *gin.Context) {
			response.Fail(ctx, errno.MethodNotAllowError)
		},
	)

	engine.GET(
		"/", func(ctx *gin.Context) {
			response.Success(
				ctx, gin.H{
					"message": "Welcome to khan",
					"version": constant.VERSION,
					"author":  "wechat:  wxksky2022",
					"remark":  "Add me, please note khan",
					"github":  "https://github.com/maogou",
					"time":    time.Now().Format(time.DateTime),
				},
			)
		},
	)
}

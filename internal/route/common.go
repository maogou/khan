package router

import (
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
			response.SuccessMsg(ctx, "欢迎使用机器人=>"+time.Now().Format(time.DateTime))
		},
	)
}

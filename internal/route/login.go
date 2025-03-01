package router

import (
	"smallBot/internal/handler/login"
	"smallBot/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func initLoginHandler(engine *gin.Engine, sdk *khan.Khan) {

	loginHandler := login.NewLoginHandler(sdk)

	engine.POST("/v1/api/login/createApp", loginHandler.CreateApp)
	engine.POST("/v1/api/login/getLoginQrCode", loginHandler.LoginQrCode)
	engine.POST("/v1/api/login/checkLogin", loginHandler.CheckLogin)
	engine.POST("/v1/api/login/setCallback", loginHandler.SetCallback)
	engine.POST("/v1/api/login/checkOnline", loginHandler.CheckOnline)
	engine.POST("/v1/api/login/reconnection", loginHandler.Reconnection)
	engine.POST("/v1/api/login/logout", loginHandler.Logout)
}

package router

import (
	"smallBot/internal/handler/login"
	"smallBot/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func initLoginHandler(engine *gin.Engine, sdk *khan.Khan) {

	loginHandler := login.NewLoginHandler(sdk)

	engine.POST("/v2/api/login/getLoginQrCode", loginHandler.LoginQrCode)
	engine.POST("/v2/api/login/checkLogin", loginHandler.CheckLogin)
	engine.POST("/v2/api/login/setCallback", loginHandler.SetCallback)
}

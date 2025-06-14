package router

import (
	"maogou/khan/internal/handler/personal"
	"maogou/khan/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func initPersonal(engine *gin.Engine, sdk *khan.Khan) {
	personalHandler := personal.NewPersonalHandler(sdk)

	engine.POST("/v1/api/personal/getQrCode", personalHandler.Qrcode)
	engine.POST("/v1/api/personal/getProfile", personalHandler.Profile)
	engine.POST("/v1/api/personal/getSafetyInfo", personalHandler.Safety)
	engine.POST("/v1/api/personal/privacySettings", personalHandler.PrivacySetting)
	engine.POST("/v1/api/personal/updateHeadImg", personalHandler.UploadHdImg)
	engine.GET("/v1/api/personal/license", personalHandler.License)
	engine.POST("/v1/api/personal/updateProfile", personalHandler.UpdateProfile)
}

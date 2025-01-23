package router

import (
	"smallBot/internal/handler/personal"
	"smallBot/internal/sdk/gewe"

	"github.com/gin-gonic/gin"
)

func initPersonal(engine *gin.Engine, sdk *gewe.Gewe) {
	personalHandler := personal.NewPersonalHandler(sdk)

	engine.POST("/api/v2/personal/getQrCode", personalHandler.Qrcode)
	engine.POST("/api/v2/personal/getProfile", personalHandler.Profile)
	engine.POST("/api/v2/personal/getSafetyInfo", personalHandler.Safety)
	engine.POST("/api/v2/personal/privacySettings", personalHandler.PrivacySetting)
	engine.POST("/api/v2/personal/updateHeadImg", personalHandler.UploadHdImg)
}

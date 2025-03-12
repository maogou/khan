package router

import (
	"github.com/gin-gonic/gin"
	"smallBot/internal/handler/setting"
	"smallBot/internal/sdk/khan"
)

func initSettingRoute(route *gin.Engine, sdk *khan.Khan) {
	settingHandler := setting.NewSettingHandler(sdk)

	route.POST("/v1/api/setting/license", settingHandler.SetLicense)
}

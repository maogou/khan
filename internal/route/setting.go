package router

import (
	"github.com/gin-gonic/gin"
	"maogou/khan/internal/handler/setting"
	"maogou/khan/internal/sdk/khan"
)

func initSettingRoute(route *gin.Engine, sdk *khan.Khan) {
	settingHandler := setting.NewSettingHandler(sdk)

	route.POST("/v1/api/setting/license", settingHandler.SetLicense)
	route.POST("/v1/api/setting/switch", settingHandler.SetSwitch)
}

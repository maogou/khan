package router

import (
	"maogou/khan/internal/handler/favor"
	"maogou/khan/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func initFavorRoute(engine *gin.Engine, sdk *khan.Khan) {
	favorHandler := favor.NewFavorHandler(sdk)

	engine.POST("/v1/api/favor/sync", favorHandler.Sync)
	engine.POST("/v1/api/favor/getContent", favorHandler.Detail)
	engine.POST("/v1/api/favor/delete", favorHandler.Delete)
}

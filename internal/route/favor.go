package router

import (
	"github.com/gin-gonic/gin"
	"smallBot/internal/handler/favor"
	"smallBot/internal/sdk/khan"
)

func initFavorRoute(engine *gin.Engine, sdk *khan.Khan) {
	favorHandler := favor.NewFavorHandler(sdk)

	engine.POST("/v2/api/favor/sync", favorHandler.Sync)
	engine.POST("/v2/api/favor/getContent", favorHandler.Detail)
	engine.POST("/v2/api/favor/delete", favorHandler.Delete)
}

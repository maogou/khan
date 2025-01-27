package router

import (
	"smallBot/internal/handler/label"
	"smallBot/internal/sdk/gewe"

	"github.com/gin-gonic/gin"
)

func initLabelRoute(engine *gin.Engine, sdk *gewe.Gewe) {

	labelHandler := label.NewLabelHandler(sdk)

	engine.POST("/v2/api/label/add", labelHandler.Add)
	engine.POST("/v2/api/label/delete", labelHandler.Delete)
	engine.POST("/v2/api/label/list", labelHandler.List)
	engine.POST("/v2/api/label/modifyMemberList", labelHandler.Modify)
}

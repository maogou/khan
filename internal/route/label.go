package router

import (
	"maogou/khan/internal/handler/label"
	"maogou/khan/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func initLabelRoute(engine *gin.Engine, sdk *khan.Khan) {

	labelHandler := label.NewLabelHandler(sdk)

	engine.POST("/v1/api/label/add", labelHandler.Add)
	engine.POST("/v1/api/label/delete", labelHandler.Delete)
	engine.POST("/v1/api/label/list", labelHandler.List)
	engine.POST("/v1/api/label/modifyMemberList", labelHandler.Modify)
}

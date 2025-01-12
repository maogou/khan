package router

import (
	"smallBot/internal/handler/callback"
	"smallBot/internal/sdk/gewe"

	"github.com/gin-gonic/gin"
)

func initCallbackRoute(route *gin.Engine, sdk *gewe.Gewe) {
	callbackHandler := callback.NewCallbackHandler(sdk)

	route.POST("/api/v1/collect", callbackHandler.Collect)
}

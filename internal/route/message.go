package router

import (
	"smallBot/internal/handler/message"
	"smallBot/internal/sdk/gewe"

	"github.com/gin-gonic/gin"
)

func initMessageRoute(route *gin.Engine, sdk *gewe.Gewe) {
	messageHandler := message.NewMessageHandler(sdk)

	route.POST("/api/v1/message/text", messageHandler.PostText)
}

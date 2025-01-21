package router

import (
	"smallBot/internal/handler/message"
	"smallBot/internal/sdk/gewe"

	"github.com/gin-gonic/gin"
)

func initMessageRoute(route *gin.Engine, sdk *gewe.Gewe) {
	messageHandler := message.NewMessageHandler(sdk)

	route.POST("/api/v2/message/postText", messageHandler.PostText)
	route.POST("/api/v2//message/postImage", messageHandler.PostImage)
	route.POST("/api/v2//message/postFile", messageHandler.PostFile)
	route.POST("/api/v2//message/postLink", messageHandler.PostLink)
}

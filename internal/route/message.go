package router

import (
	"smallBot/internal/handler/message"
	"smallBot/internal/sdk/gewe"

	"github.com/gin-gonic/gin"
)

func initMessageRoute(route *gin.Engine, sdk *gewe.Gewe) {
	messageHandler := message.NewMessageHandler(sdk)

	route.POST("/v2/api/message/postText", messageHandler.PostText)
	route.POST("/v2/api//message/postImage", messageHandler.PostImage)
	route.POST("/v2/api//message/postFile", messageHandler.PostFile)
	route.POST("/v2/api//message/postLink", messageHandler.PostLink)

	route.POST("/v2/api//message/postNameCard", messageHandler.PostNameCard)
	route.POST("/v2/api//message/postEmoji", messageHandler.PostEmoji)
}

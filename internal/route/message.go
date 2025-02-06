package router

import (
	"smallBot/internal/handler/message"
	"smallBot/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func initMessageRoute(route *gin.Engine, sdk *khan.Khan) {
	messageHandler := message.NewMessageHandler(sdk)

	route.POST("/v2/api/message/postText", messageHandler.PostText)
	route.POST("/v2/api/message/postImage", messageHandler.PostImage)
	route.POST("/v2/api/message/postFile", messageHandler.PostFile)
	route.POST("/v2/api/message/postLink", messageHandler.PostLink)
	route.POST("/v2/api/message/postVoice", messageHandler.PostVoice)
	route.POST("/v2/api/message/postNameCard", messageHandler.PostNameCard)
	route.POST("/v2/api/message/postEmoji", messageHandler.PostEmoji)
	route.POST("/v2/api/message/postVideo", messageHandler.PostVideo)
	route.POST("/v2/api/message/downloadImage", messageHandler.DownloadImg)
}

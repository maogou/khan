package router

import (
	"smallBot/internal/handler/message"
	"smallBot/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func initMessageRoute(route *gin.Engine, sdk *khan.Khan) {
	messageHandler := message.NewMessageHandler(sdk)

	route.POST("/v1/api/message/postText", messageHandler.PostText)
	route.POST("/v1/api/message/postImage", messageHandler.PostImage)
	route.POST("/v1/api/message/postFile", messageHandler.PostFile)
	route.POST("/v1/api/message/postLink", messageHandler.PostLink)
	route.POST("/v1/api/message/postVoice", messageHandler.PostVoice)
	route.POST("/v1/api/message/postNameCard", messageHandler.PostNameCard)
	route.POST("/v1/api/message/postEmoji", messageHandler.PostEmoji)
	route.POST("/v1/api/message/postVideo", messageHandler.PostVideo)
	route.POST("/v1/api/message/downloadImage", messageHandler.DownloadImg)
	route.POST("/v1/api/message/revokeMsg", messageHandler.RevokeMsg)
	route.POST("/v1/api/message/forwardMiniApp", messageHandler.ForwardMiniApp)
	route.POST("/v1/api/message/forwardUrl", messageHandler.ForwardUrl)
	route.POST("/v1/api/message/forwardVideo", messageHandler.ForwardVideo)
	route.POST("/v1/api/message/forwardImage", messageHandler.ForwardImage)
}

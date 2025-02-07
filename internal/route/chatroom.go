package router

import (
	"github.com/gin-gonic/gin"
	"smallBot/internal/handler/chatroom"
	"smallBot/internal/sdk/khan"
)

func initChatroomHandler(engine *gin.Engine, sdk *khan.Khan) {
	chatroomHandler := chatroom.NewChatRoomHandler(sdk)

	engine.POST("/v2/api/group/createChatroom", chatroomHandler.Create)
	engine.POST("/v2/api/group/modifyChatroomName", chatroomHandler.ModifyName)
	engine.POST("/v2/api/group/modifyChatroomRemark", chatroomHandler.ModifyRemark)
	engine.POST("/v2/api/group/modifyChatroomNickNameForSelf", chatroomHandler.ModifyMyselfName)
	engine.POST("/v2/api/group/setChatroomAnnouncement", chatroomHandler.SetAnnouncement)
	engine.POST("/v2/api/group/getChatroomAnnouncement", chatroomHandler.GetAnnouncement)
	engine.POST("/v2/api/group/getChatroomInfo", chatroomHandler.Detail)
	engine.POST("/v2/api/group/getChatroomMemberList", chatroomHandler.Member)
	engine.POST("/v2/api/group/pinChat", chatroomHandler.TopPing)
	engine.POST("/v2/api/group/setMsgSilence", chatroomHandler.SetSilence)
	engine.POST("/v2/api/group/getChatroomQrCode", chatroomHandler.QrCode)
}

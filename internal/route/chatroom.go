package router

import (
	"smallBot/internal/handler/chatroom"
	"smallBot/internal/sdk/khan"

	"github.com/gin-gonic/gin"
)

func initChatroomHandler(engine *gin.Engine, sdk *khan.Khan) {
	chatroomHandler := chatroom.NewChatRoomHandler(sdk)

	engine.POST("/v1/api/group/createChatroom", chatroomHandler.Create)
	engine.POST("/v1/api/group/modifyChatroomName", chatroomHandler.ModifyName)
	engine.POST("/v1/api/group/modifyChatroomRemark", chatroomHandler.ModifyRemark)
	engine.POST("/v1/api/group/modifyChatroomNickNameForSelf", chatroomHandler.ModifyMyselfName)
	engine.POST("/v1/api/group/setChatroomAnnouncement", chatroomHandler.SetAnnouncement)
	engine.POST("/v1/api/group/getChatroomAnnouncement", chatroomHandler.GetAnnouncement)
	engine.POST("/v1/api/group/getChatroomInfo", chatroomHandler.Detail)
	engine.POST("/v1/api/group/getChatroomMemberList", chatroomHandler.Member)
	engine.POST("/v1/api/group/pinChat", chatroomHandler.TopPing)
	engine.POST("/v1/api/group/setMsgSilence", chatroomHandler.SetSilence)
	engine.POST("/v1/api/group/getChatroomQrCode", chatroomHandler.QrCode)
	engine.POST("/v1/api/group/removeMember", chatroomHandler.DeleteMember)
	engine.POST("/v1/api/group/inviteMember", chatroomHandler.InviteMember)
	engine.POST("/v1/api/group/adminOperate", chatroomHandler.AdminOperate)
	engine.POST("/v1/api/group/roomAccessApplyCheckApprove", chatroomHandler.ConfirmInvite)
	engine.POST("/v1/api/group/joinRoomUsingQRCode", chatroomHandler.ScanQrcodeEnter)
}

package chatroom

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/chatroom"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *ChatRoomHandler) SetAnnouncement(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->Announcement方法")

	var req v1.ChatroomSetAnnouncementRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->Announcement方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ChatroomSetAnnouncement(
		ctx, chatroom.ChatroomSetAnnouncementRequest{
			Appid:   req.AppId,
			GroupId: req.ChatroomId,
			Content: req.Content,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomSetAnnouncement方法失败")
		response.Fail(ctx, errno.ChatroomSetAnnouncementError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret !=0 ->调用ChatroomSetAnnouncement方法失败")
		response.Fail(ctx, errno.ChatroomSetAnnouncementError)
		return
	}

	response.SuccessMsg(ctx, "修改群聊公告成功")
}

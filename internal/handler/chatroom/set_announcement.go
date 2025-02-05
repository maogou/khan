package chatroom

import (
	"github.com/gin-gonic/gin"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
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
		ctx, transform.ChatroomSetAnnouncementRequest{
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

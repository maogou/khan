package chatroom

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/chatroom"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *ChatRoomHandler) GetAnnouncement(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->GetAnnouncement方法")

	var req v1.ChatroomGetAnnouncementRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->GetAnnouncement方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ChatroomAnnouncement(
		ctx, chatroom.ChatroomAnnouncementRequest{
			Appid:   req.AppId,
			GroupId: req.ChatroomId,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomAnnouncement方法失败")
		response.Fail(ctx, errno.ChatroomGetAnnouncementError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomAnnouncement方法失败")
		response.Fail(ctx, errno.ChatroomGetAnnouncementError)
		return
	}

	response.Success(
		ctx, v1.ChatroomGetAnnouncementResponse{
			Announcement:       resp.Data.Announcement,
			AnnouncementEditor: resp.Data.AnnouncementEditor,
			PublishTime:        resp.Data.AnnouncementPublishTime,
		},
	)
}

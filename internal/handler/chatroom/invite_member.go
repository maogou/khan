package chatroom

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/chatroom"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func (c *ChatRoomHandler) InviteMember(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->InviteMember方法")

	var req v1.ChatroomInviteMemberRequset

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->InviteMember方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	wxids := strings.Trim(req.Wxids, ",")

	resp, err := c.sdk.ChatroomInvite(
		ctx, chatroom.ChatroomInviteRequest{
			Appid:   req.AppId,
			GroupId: req.ChatroomId,
			ToWxid:  strings.Split(wxids, ","),
			Cause:   req.Reason,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->InviteMember方法调用失败")
		response.Fail(ctx, errno.ChatroomInviteMemberError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->InviteMember方法调用失败")
		response.Fail(ctx, errno.ChatroomInviteMemberError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Any(
			"msg_err", resp.Data.BaseResponse.ErrMsg,
		).Msg("BaseResponse.Ret !=0 ->调用ChatroomInvite方法失败")
		response.Fail(ctx, errno.ChatroomInviteMemberError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")
}

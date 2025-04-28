package chatroom

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/chatroom"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *ChatRoomHandler) AddGroupMemberAsFriend(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->AddGroupMemberAsFriend方法")

	var req v1.ChatroomAddGroupMemberAsFriendRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->AddGroupMemberAsFriend方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ChatroomAddMemberAsFriend(
		ctx, chatroom.ChatroomAddMemberAsFriendRequest{
			Appid:   req.AppId,
			GroupId: req.ChatroomId,
			Content: req.Content,
			ToWxid:  req.MemberWxid,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomAddMemberAsFriend方法失败")
		response.Fail(ctx, errno.ChatroomAddMemberAsFriendError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("调用ChatroomAddMemberAsFriend方法resp.Ret != 0失败")
		response.Fail(ctx, errno.ChatroomAddMemberAsFriendError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Msg("调用ChatroomAddMemberAsFriend方法resp.Data.BaseResponse.Ret != 0失败")
		response.Fail(ctx, errno.ChatroomAddMemberAsFriendError)
		return
	}

	response.Success(
		ctx, v1.ChatroomAddGroupMemberAsFriendResponse{
			V3: resp.Data.V1,
		},
	)
}

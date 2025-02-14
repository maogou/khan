package chatroom

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/chatroom"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *ChatRoomHandler) Member(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->Member方法")

	var (
		req    v1.ChatroomMemberRequest
		result = v1.ChatroomMemberResponse{
			AdminWxid:  make([]string, 0),
			MemberList: make([]v1.ChatroomMemberItem, 0),
		}
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->Member方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ChatroomMember(
		ctx, chatroom.ChatroomMemberRequest{
			Appid:   req.AppId,
			GroupId: req.ChatroomId,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomMember方法失败")
		response.Fail(ctx, errno.ChatroomMemberError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomMember方法失败")
		response.Fail(ctx, errno.ChatroomMemberError)
		return
	}

	result.ChatroomOwner = resp.Data.ChatRoomOwner
	result.AdminWxid = make([]string, 0)

	if len(resp.Data.AdminUserNameList) > 0 {
		for _, item := range resp.Data.AdminUserNameList {
			result.AdminWxid = append(result.AdminWxid, item.String)
		}
	}

	if len(resp.Data.NewChatroomData.ChatRoomMember) > 0 {
		for _, item := range resp.Data.NewChatroomData.ChatRoomMember {
			result.MemberList = append(
				result.MemberList, v1.ChatroomMemberItem{
					Wxid:            item.UserName,
					NickName:        item.NickName,
					InviterUserName: item.InviterUserName,
					MemberFlag:      item.ChatroomMemberFlag,
					DisplayName:     item.DisplayName,
					SmallHeadImgUrl: item.SmallHeadImgUrl,
					BigHeadImgUrl:   item.BigHeadImgUrl,
				},
			)
		}
	}

	response.Success(ctx, result)
}

package chatroom

import (
	"github.com/gin-gonic/gin"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/chatroom"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
)

func (c *ChatRoomHandler) Detail(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->Info方法")

	var (
		req    v1.ChatroomDetailRequest
		result = v1.ChatroomDetailResponse{
			MemberList: make([]v1.ChatroomMemberItem, 0),
		}
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->Info方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ChatroomDetail(
		ctx, chatroom.ChatroomDetailRequest{
			Appid:       req.AppId,
			GroupIdList: []string{req.ChatroomId},
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomDetail方法失败")
		response.Fail(ctx, errno.ChatroomDetailError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomDetail方法失败")
		response.Fail(ctx, errno.ChatroomDetailError)
		return
	}

	if len(resp.Data.ContactList) == 0 || len(resp.Data.ContactList[0].NewChatroomData.ChatRoomMember) == 0 {
		response.Success(ctx, result)
		return
	}

	result.ChatroomId = req.ChatroomId
	result.UserName = resp.Data.ContactList[0].UserName.String
	result.NickName = resp.Data.ContactList[0].NickName.String
	result.PyInitial = resp.Data.ContactList[0].Pyinitial.String
	result.QuanPin = resp.Data.ContactList[0].QuanPin.String
	result.Sex = resp.Data.ContactList[0].Sex
	result.Remark = resp.Data.ContactList[0].Remark.String
	result.RemarkPyInitial = resp.Data.ContactList[0].RemarkPyinitial.String
	result.RemarkQuanPin = resp.Data.ContactList[0].RemarkQuanPin.String
	result.ChatRoomNotify = resp.Data.ContactList[0].ChatRoomNotify
	result.ChatRoomOwner = resp.Data.ContactList[0].ChatRoomOwner
	result.SmallHeadImgUrl = resp.Data.ContactList[0].SmallHeadImgUrl

	for _, member := range resp.Data.ContactList[0].NewChatroomData.ChatRoomMember {
		result.MemberList = append(
			result.MemberList, v1.ChatroomMemberItem{
				Wxid:            member.UserName,
				NickName:        member.NickName,
				InviterUserName: member.InviterUserName,
				MemberFlag:      member.ChatroomMemberFlag,
				DisplayName:     "",
				BigHeadImgUrl:   "",
				SmallHeadImgUrl: "",
			},
		)
	}

	response.Success(ctx, result)
}

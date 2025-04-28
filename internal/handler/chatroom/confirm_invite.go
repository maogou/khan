package chatroom

import (
	"encoding/xml"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/chatroom"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

type Sysmsg struct {
	XMLName xml.Name `xml:"sysmsg"`
	Content struct {
		Link     Link   `xml:"link"`
		RoomName string `xml:"RoomName"`
	} `xml:"NewXmlChatRoomAccessVerifyApplication"`
}

type Link struct {
	Ticket          string `xml:"ticket"`
	InviterUsername string `xml:"inviterusername"`
	MemberList      struct {
		Members []Member `xml:"member"`
	} `xml:"memberlist"`
}

type Member struct {
	Username string `xml:"username"`
}

func (c *ChatRoomHandler) ConfirmInvite(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->ConfirmInvite方法")

	var (
		req     v1.ChatroomConfimInviteRequest
		sysmsg  Sysmsg
		members []string
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->ConfirmInvite方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	content := strings.NewReplacer(
		"\t", "",
		"\n", "",
	).Replace(req.MsgContent)

	if err := xml.Unmarshal([]byte(content), &sysmsg); err != nil {
		log.C(ctx).Error().Err(err).Msg("xml解析49类型的消息失败")
		response.Fail(ctx, errno.ValidateError)
		return

	}

	for _, member := range sysmsg.Content.Link.MemberList.Members {
		members = append(members, member.Username)
	}

	resp, err := c.sdk.AccessApplyCheckApprove(
		ctx, chatroom.ChatroomConfirInviteRequest{
			Appid:       req.AppId,
			GroupId:     req.ChatroomId,
			InviterWxid: sysmsg.Content.Link.InviterUsername,
			Msgid:       req.NewMsgId,
			Ticket:      sysmsg.Content.Link.Ticket,
			ToWxidList:  members,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用sdk->AccessApplyCheckApprove方法失败")
		response.Fail(ctx, errno.ChatroomConfirInviteError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用sdk->AccessApplyCheckApprove方法失败")
		response.Fail(ctx, errno.ChatroomConfirInviteError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("BaseResponse.Ret != 0 ->调用sdk->AccessApplyCheckApprove方法失败")
		response.Fail(ctx, errno.ChatroomConfirInviteError)
		return
	}

	response.SuccessMsg(ctx, "确认成功")
}

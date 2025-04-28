package chatroom

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/chatroom"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func (c *ChatRoomHandler) DeleteMember(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->DeleteMember方法")

	var req v1.ChatroomDeleteMemberRequset

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->DeleteMember方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	wids := strings.Trim(req.Wxids, ",")

	resp, err := c.sdk.ChatroomDeleteMember(
		ctx, chatroom.ChatroomDeleteMebmerRequest{
			Appid:      req.AppId,
			GroupId:    req.ChatroomId,
			ToWxidList: strings.Split(wids, ","),
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->DeleteMember方法调用失败")
		response.Fail(ctx, errno.ChatroomDeleteMemberError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("ChatRoomHandler->DeleteMember方法resp.Ret != 0失败")
		response.Fail(ctx, errno.ChatroomDeleteMemberError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Msg("ChatRoomHandler->DeleteMember方法resp.Data.BaseResponse.Ret != 0失败")
		response.Fail(ctx, errno.ChatroomDeleteMemberError)
		return
	}

	response.SuccessMsg(ctx, "删除成功")
}

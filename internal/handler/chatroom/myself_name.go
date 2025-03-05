package chatroom

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/chatroom"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *ChatRoomHandler) ModifyMyselfName(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->ModifyMyselfName方法")

	var req v1.ChatroomModifyMyselfNameRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->ModifyMyselfName方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ChatroomModifyMyNickName(
		ctx, chatroom.ChatroomModifyNickNameRequest{
			Appid:    req.AppId,
			GroupId:  req.ChatroomId,
			NickName: req.NickName,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomModifyMyNickName方法失败")
		response.Fail(ctx, errno.ChatroomModifyMyselfNameError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用ChatroomModifyMyNickName方法失败")
		response.Fail(ctx, errno.ChatroomModifyMyselfNameError)
		return
	}

	response.SuccessMsg(ctx, "修改自己在群聊昵称成功")
}

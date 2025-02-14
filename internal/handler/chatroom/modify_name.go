package chatroom

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/chatroom"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *ChatRoomHandler) ModifyName(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->ModifyName方法")

	var req v1.ChatroomModifyNameRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->ModifyName方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ChatroomModifyName(
		ctx, chatroom.ChatroomModifyNameRequest{
			Appid:   req.AppId,
			GroupId: req.ChatroomId,
			Name:    req.ChatroomName,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomModifyName方法失败")
		response.Fail(ctx, errno.ChatroomModifyNameError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用ChatroomModifyName方法失败")
		response.Fail(ctx, errno.ChatroomModifyNameError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")
}

package chatroom

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/chatroom"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *ChatRoomHandler) MoveContractList(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->MoveContractList方法")

	var req v1.ChatroomMoveContractListRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->MoveContractList方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ChatroomMoveContract(
		ctx, chatroom.ChatroomMoveContractRequest{
			Appid:   req.AppId,
			GroupId: req.ChatroomId,
			Value:   req.OperType,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用MoveContractList方法失败")
		response.Fail(ctx, errno.ChatroomMoveContractListError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用MoveContractList方法失败")
		response.Fail(ctx, errno.ChatroomMoveContractListError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")
}

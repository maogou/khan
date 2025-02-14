package chatroom

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/chatroom"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func (c *ChatRoomHandler) SetSilence(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->SetSilence方法")

	var req v1.ChatroomSetSilenceRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->SetSilence方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	isHasChatroom := strings.HasSuffix(req.ChatroomId, "@chatroom")

	value := lo.TernaryF(
		isHasChatroom, func() int {
			return lo.Ternary(req.Silence, 0, 1)
		}, func() int {
			return lo.Ternary(req.Silence, 1, 0)
		},
	)

	resp, err := c.sdk.ChatroomSetMsgSilence(
		ctx, chatroom.ChatroomSetSilenceRequest{
			Appid:   req.AppId,
			GroupId: req.ChatroomId,
			Value:   value,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomSetMsgSilence方法失败")
		response.Fail(ctx, errno.ChatroomSetSilenceError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomSetMsgSilence方法失败")
		response.Fail(ctx, errno.ChatroomSetSilenceError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")
}

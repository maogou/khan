package chatroom

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/chatroom"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func (c *ChatRoomHandler) TopPing(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->TopPing方法")

	var req v1.ChatroomTopPingRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->TopPing方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	isHasChatroom := strings.HasSuffix(req.ChatroomId, "@chatroom")

	value := lo.TernaryF(
		isHasChatroom, func() int {
			return lo.Ternary(req.Top, 2050, 2)
		}, func() int {
			return lo.Ternary(req.Top, 34823, 32775)
		},
	)

	resp, err := c.sdk.ChatroomTopPing(
		ctx, chatroom.ChatroomTopPingRequest{
			Appid:   req.AppId,
			GroupId: req.ChatroomId,
			Value:   value,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->TopPing方法调用失败")
		response.Fail(ctx, errno.ChatroomTopPingError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->TopPing方法调用失败")
		response.Fail(ctx, errno.ChatroomTopPingError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")
}

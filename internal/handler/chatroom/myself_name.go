package chatroom

import (
	"github.com/gin-gonic/gin"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
)

func (c *ChatRoomHandler) ModifyMyselfName(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->ModifyMyselfName方法")

	var req v1.ChatroomModifyNameRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->ModifyMyselfName方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ChatroomModifyMyNickName(
		ctx, transform.ChatroomModifyNickNameRequest{
			Appid:    req.AppId,
			GroupId:  req.ChatroomId,
			NickName: req.ChatroomName,
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

	response.SuccessMsg(ctx, "修改群聊名称成功")
}

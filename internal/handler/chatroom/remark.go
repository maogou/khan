package chatroom

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/chatroom"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *ChatRoomHandler) ModifyRemark(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->Remark方法")

	var req v1.ChatroomSetRemarkRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ChatroomModifyRemark(
		ctx, chatroom.ChatroomModifyRemarkRequest{
			Appid:   req.AppId,
			GroupId: req.ChatroomId,
			Name:    req.ChatroomRemark,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomModifyRemark方法失败")
		response.Fail(ctx, errno.ChatroomModifyRemarkError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用ChatroomModifyRemark方法失败")
		response.Fail(ctx, errno.ChatroomModifyRemarkError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")

}

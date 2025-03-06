package chatroom

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/chatroom"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *ChatRoomHandler) AdminOperate(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->AdminOperate方法")

	var req v1.ChatroomAdminOperateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->AdminOperate方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ChatroomAdminOperate(
		ctx, chatroom.ChatroomAdminOperateRequest{
			Appid:      req.AppId,
			GroupId:    req.ChatroomId,
			Value:      *req.OperType,
			ToWxidList: req.Wxids,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomAdminOperate方法失败")
		response.Fail(ctx, errno.ChatroomAdminOperateError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("ret != 0 ->调用ChatroomAdminOperate方法失败")
		response.Fail(ctx, errno.ChatroomAdminOperateError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Any(
			"msg_err", resp.Data.BaseResponse.ErrMsg,
		).Msg("BaseResponse.Ret !=0 ->调用ChatroomAdminOperate方法失败")
		response.Fail(ctx, errno.ChatroomAdminOperateError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")
}

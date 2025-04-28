package chatroom

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/chatroom"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *ChatRoomHandler) QrCode(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ChatRoomHandler->QrCode方法")

	var req v1.ChatroomQrcodeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ChatroomQrcode(
		ctx, chatroom.ChatroomQrcodeRequest{
			AppId:   req.AppId,
			GroupId: req.ChatroomId,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomQrcode方法失败")
		response.Fail(ctx, errno.ChatroomQrcodeError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用ChatroomQrcode方法失败")
		response.Fail(ctx, errno.ChatroomQrcodeError)
		return
	}

	response.Success(
		ctx, v1.ChatroomQrcodeResponse{
			QrBase64: "data:image/jpg;base64," + resp.Data.Qrcode.Buffer,
			QrTips:   resp.Data.RevokeQrcodeWording,
		},
	)
}

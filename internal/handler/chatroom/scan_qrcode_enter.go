package chatroom

import (
	"github.com/gin-gonic/gin"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/chatroom"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
)

func (c *ChatRoomHandler) ScanQrcodeEnter(ctx *gin.Context) {
	log.C(ctx).Info().Msg("开始调用ChatRoomHandler->ScanQrcodeEnter方法")

	var req v1.ChatroomScanQrcodeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("ChatRoomHandler->ScanQrcodeEnter方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ScanQrcodeEnter(
		ctx, chatroom.ChatroomScanQrcodeEnterRequest{
			Appid: req.AppId,
			Config: chatroom.ChatroomScanQrcodeEnterConfig{
				Type: 0,
				Url:  req.QrUrl,
			},
			Scene: 10001,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ScanQrcodeEnter方法失败")
		response.Fail(ctx, errno.ChatroomScanQrcodeEnterError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("调用ScanQrcodeEnter方法resp.Ret != 0失败")
		response.Fail(ctx, errno.ChatroomScanQrcodeEnterError)
		return
	}

	response.Success(
		ctx, v1.ChatroomScanQrcodeResponse{
			ChatroomId:   resp.Data.Result.Qid,
			ChatroomName: resp.Data.Result.GroupName,
		},
	)
}

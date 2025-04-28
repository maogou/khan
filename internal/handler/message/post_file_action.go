package message

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/message"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (m *MessageHandler) PostFile(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->PostFile方法")

	var req v1.PostFileRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := m.sdk.PostFile(
		ctx, message.PostFileRequest{
			Appid: req.AppId,
			Config: message.PostFileConfig{
				FileName: req.FileName,
			},
			FileLink: req.FileUrl,
			ToWxid:   req.ToWxid,
		},
	)

	if err != nil {
		log.C(ctx).Error().Msg("发送文件失败")
		response.Fail(ctx, errno.PostFileError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用PostFile方法发送消息失败")
		response.Fail(ctx, errno.PostFileError)
		return
	}

	response.Success(
		ctx, v1.PostFileResponse{
			ToWxid:     req.ToWxid,
			CreateTime: resp.Data.SendMsgStatus.CreateTime,
			MsgId:      resp.Data.SendMsgStatus.MsgId,
			NewMsgId:   resp.Data.SendMsgStatus.NewMsgId,
			Type:       resp.Data.SendMsgStatus.Type,
		},
	)

}

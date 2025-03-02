package message

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/message"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (m *MessageHandler) PostVoice(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->PostVoice方法")

	var req v1.PostVoiceRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := m.sdk.PostVoice(
		ctx, message.PostVoiceRequest{
			Appid:    req.AppId,
			ToWxid:   req.ToWxid,
			FileLink: req.VoiceUrl,
			Second:   req.VoiceDuration / 1000,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用PostVoice方法发送消息失败")
		response.Fail(ctx, errno.PostVoiceError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Any("resp", resp).Msg("调用PostVoice方法发送消息失败")
		response.Fail(ctx, errno.PostVoiceError)
		return
	}

	response.Success(
		ctx, v1.PostVoiceResponse{
			ToWxid:     req.ToWxid,
			MsgId:      resp.Data.MsgId,
			NewMsgId:   resp.Data.NewMsgId,
			CreateTime: resp.Data.CreateTime,
			Type:       34,
		},
	)
}

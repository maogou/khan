package message

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/message"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/help"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *MessageHandler) PostVideo(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->PostVideo方法")

	var req v1.PostVideoRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	base64, err := help.GetFileBase64(m.sdk.Client(), req.ThumbUrl)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.GetFileBase64Error)
		return
	}

	resp, err := m.sdk.PostVideo(
		ctx, message.PostVideoRequest{
			Appid:  req.AppId,
			ToWxid: req.ToWxid,
			Config: message.PostVideoConfig{
				VideoPlayLength: req.VideoDuration,
				VideoThumbnail:  base64,
			},
			FileLink: req.VideoUrl,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用PostVideo方法发送消息失败")
		response.Fail(ctx, errno.PostVideoError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Any("resp", resp).Msg("调用PostVideo方法发送消息失败")
		response.Fail(ctx, errno.PostVideoError)
		return
	}

	response.Success(
		ctx, v1.PostVideoResponse{
			ToWxid:     req.ToWxid,
			MsgId:      resp.Data.SendMsgStatus.MsgId,
			NewMsgId:   resp.Data.SendMsgStatus.NewMsgId,
			CreateTime: time.Now().Local().Unix(),
			Type:       43,
			FileId:     resp.Data.OtherStatus.FileId,
			AesKey:     resp.Data.SendMsgStatus.Aeskey,
			Length:     req.VideoDuration,
		},
	)
}

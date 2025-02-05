package message

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
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

	resp, err := m.sdk.PostVideo(
		ctx, transform.PostVideoRequest{
			Appid:  req.AppId,
			ToWxid: req.ToWxid,
			Config: transform.PostVideoConfig{
				VideoPlayLength: req.VideoDuration,
				VideoThumbnail:  req.ThumbUrl,
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

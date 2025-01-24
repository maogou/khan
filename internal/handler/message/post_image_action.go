package message

import (
	v1 "smallBot/api/gewe/v1"
	"smallBot/api/gewe/v1/transform"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (m *MessageHandler) PostImage(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->PostImage方法")

	var req v1.PostImageRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := m.sdk.PostImage(
		ctx, transform.PostImage{
			AppId:    req.AppId,
			Config:   transform.PostImageConfig{Thumbnail: true},
			FileLink: req.ImgUrl,
			ToWxid:   req.ToWxid,
		},
	)
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用PostImage方法发送消息失败")
		response.Fail(ctx, errno.PostImageError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用PostImage方法发送消息失败")
		response.Fail(ctx, errno.PostImageError)
		return
	}

	log.C(ctx).Info().Any("req", req).Any("resp", resp).Msg("发送图片成功")

	response.Success(
		ctx, v1.PostImageResponse{
			ToWxid:     req.ToWxid,
			CreateTime: resp.Data.SendMsgStatus.CreateTime,
			MsgId:      resp.Data.SendMsgStatus.Msgid,
			NewMsgId:   resp.Data.SendMsgStatus.Newmsgid,
			Type:       resp.Data.SendMsgStatus.MsgSource,
			AesKey:     resp.Data.SendMsgStatus.Aeskey,
			FileId:     resp.Data.SendMsgStatus.Fileid,
			Height:     resp.Data.BasicData.Height,
			Length:     resp.Data.BasicData.Length,
			Width:      resp.Data.BasicData.Width,
			Md5:        resp.Data.BasicData.FileMd5,
		},
	)
}

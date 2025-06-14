package message

import (
	"context"
	"encoding/base64"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/message"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (m *MessageHandler) PostLink(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->PostLink方法")

	var req v1.PostLinkRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	imgByte, err := m.downloadFile(ctx, req.ThumbUrl)
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("下载图片失败")
		response.Fail(ctx, errno.DownloadImgError)
		return
	}

	resp, err := m.sdk.PostLink(
		ctx, message.PostLinkRequest{
			Appid:  req.AppId,
			Base64: base64.StdEncoding.EncodeToString(imgByte),
			Des:    req.Desc,
			Title:  req.Title,
			ToWxid: req.ToWxid,
			Url:    req.LinkUrl,
		},
	)
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用PostLink方法发送消息失败")
		response.Fail(ctx, errno.PostLinkError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Any("resp", resp).Msg("调用PostLink方法发送消息失败")
		response.Fail(ctx, errno.PostLinkError)
		return
	}

	log.C(ctx).Info().Any("req", req).Any("resp", resp).Msg("发送消息成功")

	response.Success(
		ctx, v1.PostLinkResponse{
			ToWxid:     req.ToWxid,
			CreateTime: resp.Data.CreateTime,
			MsgId:      resp.Data.MsgId,
			NewMsgId:   resp.Data.NewMsgId,
			Type:       resp.Data.Type,
		},
	)
}

func (m *MessageHandler) downloadFile(ctx context.Context, url string) ([]byte, error) {
	log.C(ctx).Info().Msg("调用MessageHandler->downloadFile方法")

	resp, err := m.sdk.Client().R().Get(url)
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("下载文件失败")
		return nil, errno.DownloadImgError
	}

	return resp.Body(), nil
}

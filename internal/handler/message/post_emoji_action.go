package message

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/message"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *MessageHandler) PostEmoji(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->PostEmoji方法")

	var req v1.PostEmojiRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := m.sdk.PostEmoji(
		ctx, message.PostEmojiRequest{
			Appid:      req.AppId,
			Md5:        req.EmojiMd5,
			ToWxidList: []string{req.ToWxid},
			TotalLen:   req.EmojiSize,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用PostEmoji方法发送消息失败")
		response.Fail(ctx, errno.PostEmojiError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用PostEmoji方法发送消息失败")
		response.Fail(ctx, errno.PostEmojiError)
		return
	}

	response.Success(
		ctx, v1.PostEmojiResponse{
			ToWxid:     req.ToWxid,
			CreateTime: time.Now().Local().Unix(),
			MsgId:      resp.Data.EmojiItem[0].MsgId,
			NewMsgId:   resp.Data.EmojiItem[0].NewMsgId,
			Type:       47,
		},
	)
}

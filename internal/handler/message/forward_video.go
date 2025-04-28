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

func (m *MessageHandler) ForwardVideo(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->ForwardVideo方法")

	var req v1.ForwardVideoRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := m.sdk.ForwardVideo(
		ctx, message.ForwardVideoRequest{
			Appid:  req.AppId,
			ToWxid: req.ToWxid,
			Xml:    req.Xml,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ForwardVideo方法失败")
		response.Fail(ctx, errno.ForwardVideoError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("调用ForwardVideo方法resp.Ret != 0失败")
		response.Fail(ctx, errno.ForwardVideoError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Msg("调用ForwardVideo方法resp.Data.BaseResponse.Ret != 0失败")
		response.Fail(ctx, errno.ForwardVideoError)
		return
	}

	response.Success(
		ctx, v1.ForwardVideoResponse{
			ToWxid:     req.ToWxid,
			CreateTime: time.Now().Unix(),
			MsgId:      resp.Data.MsgId,
			NewMsgId:   resp.Data.NewMsgId,
			AesKey:     resp.Data.Aeskey,
		},
	)

}

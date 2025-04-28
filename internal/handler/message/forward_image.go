package message

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/message"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (m *MessageHandler) ForwardImage(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->ForwardImage方法")

	var req v1.ForwardImageRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := m.sdk.ForwardImage(
		ctx, message.ForwardImageRequest{
			Appid:  req.AppId,
			ToWxid: req.ToWxid,
			Xml:    req.Xml,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ForwardImage方法失败")
		response.Fail(ctx, errno.InternalServerError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("调用ForwardImage方法resp.Ret != 0失败")
		response.Fail(ctx, errno.InternalServerError)
		return
	}
	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Msg("调用ForwardImage方法resp.Data.BaseResponse.Ret != 0失败")
		response.Fail(ctx, errno.InternalServerError)
		return
	}

	response.Success(
		ctx, v1.ForwardImageResponse{
			ToWxid:     req.ToWxid,
			CreateTime: resp.Data.CreateTime,
			MsgId:      resp.Data.Msgid,
			NewMsgId:   resp.Data.Newmsgid,
			AesKey:     resp.Data.Aeskey,
			FileId:     resp.Data.Fileid,
		},
	)
}

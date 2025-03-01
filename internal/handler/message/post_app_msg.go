package message

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/message"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (m *MessageHandler) PostAppMsg(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->PostAppMsg方法")

	var req v1.PostAppMsgRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := m.sdk.PostAppMsg(
		ctx, message.PostAppMsgRequest{
			Appid:  req.AppId,
			ToWxid: req.ToWxid,
			Xml:    req.Appmsg,
			Type:   0,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用PostAppMsg方法失败")
		response.Fail(ctx, errno.SendAppMsgError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("调用PostAppMsg方法resp.Ret != 0失败")
		response.Fail(ctx, errno.SendAppMsgError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Msg("调用PostAppMsg方法resp.Data.BaseResponse.Ret != 0失败")
		response.Fail(ctx, errno.SendAppMsgError)
		return
	}

	response.Success(
		ctx, v1.PostAppMsgResponse{
			CreateTime: resp.Data.CreateTime,
			MsgId:      resp.Data.MsgId,
			NewMsgId:   resp.Data.NewMsgId,
			ToWxid:     resp.Data.ToUserName,
			Type:       resp.Data.Type,
		},
	)

}

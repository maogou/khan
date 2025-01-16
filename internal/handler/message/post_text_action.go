package message

import (
	v1 "smallBot/api/gewe/v1"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/help"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (m *MessageHandler) PostText(ctx *gin.Context) {
	log := help.GetQidLog(ctx)
	log.Info().Msg("调用MessageHandler->PostText方法")

	var req v1.PostTextRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, err)
		return
	}

	resp, err := m.sdk.PostText(ctx, v1.PostTextRequest{AppId: req.AppId, Content: req.Content, ToWxid: req.ToWxid})

	if err != nil {
		log.Error().Err(err).Msg("调用PostText方法发送消息失败")
		response.Fail(ctx, errno.SendMsgError)
		return
	}

	log.Info().Any("req", req).Any("resp", resp).Msg("发送消息成功")

	response.SuccessMsg(ctx, "success")

	return
}

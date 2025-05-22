package message

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (m *MessageHandler) PostText(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->PostText方法")

	var req v1.PostTextRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := m.sdk.PostText(ctx, req)
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用PostText方法发送消息失败")
		response.Fail(ctx, errno.PostTextError)
		return
	}

	log.C(ctx).Info().Any("req", req).Any("resp", resp).Msg("发送消息成功")

	response.Success(ctx, resp)
}

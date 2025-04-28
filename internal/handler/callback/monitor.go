package callback

import (
	"github.com/gin-gonic/gin"
	"maogou/khan/api/khan/v1/transform/message"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"
)

func (c *CallbackHandler) Monitor(ctx *gin.Context) {
	log.C(ctx).Info().Msg("开始调用CallbackHandler-Monitor")

	var req message.CallbackRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("CallbackHandler-Monitor参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	log.C(ctx).Info().Any("req", req).Msg("接收到第三者原始的回调数据")

	response.SuccessMsg(ctx, "success")
}

package callback

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *CallbackHandler) Collect(ctx *gin.Context) {
	log.C(ctx).Info().Msg("开始调用CallbackHandler-Collect")
	var req v1.CollectRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("CallbackHandler-Collect参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	log.C(ctx).Info().Any("req", req).Msg("接收到原始的回调数据")

	//中间件种和handler种必须使用ctx的副本
	cCp := ctx.Copy()
	go func() {
		c.chain.HandlerRequest(cCp, req)
	}()

	response.SuccessMsg(ctx, "success")
}

package callback

import (
	v1 "smallBot/api/gewe/v1"
	"smallBot/internal/pkg/help"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *CallbackHandler) Collect(ctx *gin.Context) {
	log := help.GetQidLog(ctx)
	log.Info().Msg("调用Collect方法")

	var req v1.CollectRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, err)
		return
	}

	//中间件种和handler种必须使用ctx的副本
	cCp := ctx.Copy()
	go func() {
		c.chain.HandlerRequest(cCp, req)
	}()

	log.Info().Any("req", req).Msg("收到的消息")

	response.SuccessMsg(ctx, "success")

	return
}

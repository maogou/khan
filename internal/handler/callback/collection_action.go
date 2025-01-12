package callback

import (
	"io"
	v1 "smallBot/api/gewe/v1"
	"smallBot/internal/pkg/help"
	"smallBot/internal/pkg/response"
	"smallBot/internal/pkg/validator"

	"github.com/gin-gonic/gin"
)

func (c *CallbackHandler) Collect(ctx *gin.Context) {
	log := help.GetQidLog(ctx)
	log.Info().Msg("调用Collect方法")

	var req v1.CollectRequest

	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		log.Error().Err(err).Msg("读取body失败")
		response.Fail(ctx, err)
		return
	}

	log.Info().Str("body", string(body)).Msg("接收到的body数据")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("参数验证失败")
		response.ValidatorErr(ctx, err, validator.Translate(err))
		return
	}

	log.Info().Any("req", req).Msg("收到的消息")

	response.Success(ctx, "ok")

	return
}

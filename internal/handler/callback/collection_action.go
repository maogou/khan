package callback

import (
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

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidatorErr(ctx, err, validator.Translate(err))
		return
	}

	log.Info().Any("req", req).Msg("收到的消息")

	response.Success(ctx, "ok")

	return
}

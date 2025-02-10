package transfer

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (t *TransferHandler) Import(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用TransferHandler->Import方法")

	var (
		req    v1.TransferImportRequest
		result v1.TransferImportResponse
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	rdb := t.sdk.Rdb()
	for _, v := range req.Data {
		if err := rdb.Set(ctx, v.Key, v.Value, 0).Err(); err != nil {
			result.Fail++
			log.C(ctx).Error().Err(err).Str("key", v.Key).Str("value", v.Value).Msg("redis设置数据失败")
			continue
		}

		result.Success++
	}

	response.Success(ctx, result)
}

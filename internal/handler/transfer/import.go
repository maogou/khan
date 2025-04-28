package transfer

import (
	"encoding/json"
	"fmt"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/help"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (t *TransferHandler) Import(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用TransferHandler->Import方法")

	var (
		req    v1.TransferImportRequest
		backup v1.TransferResponse
		result v1.TransferImportResponse
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	decrypted, err := help.AesDecrypt(req.Backup, []byte(constant.AesBackup))

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("解密失败")
		response.Fail(ctx, errno.AesDecryptError)
		return
	}

	fmt.Println(decrypted)

	if err = json.Unmarshal([]byte(decrypted), &backup); err != nil {
		log.C(ctx).Error().Err(err).Msg("json解析失败")
		response.Fail(ctx, errno.JsonDecodeError)
		return
	}

	rdb := t.sdk.Rdb()
	for _, v := range backup.Data {
		if err = rdb.Set(ctx, v.Key, v.Value, 0).Err(); err != nil {
			result.Fail++
			log.C(ctx).Error().Err(err).Str("key", v.Key).Str("value", v.Value).Msg("redis设置数据失败")
			continue
		}

		result.Success++
	}

	response.Success(ctx, result)
}

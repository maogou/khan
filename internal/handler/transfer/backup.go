package transfer

import (
	"encoding/json"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/help"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (t *TransferHandler) Backup(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用TransferHandler->Backup方法")

	var result = v1.TransferResponse{
		Data: make([]v1.TransferBackupItem, 0),
	}

	keys, err := t.sdk.Rdb().Keys(ctx, "wx_*").Result()

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("获取redis数据失败")
		response.Fail(ctx, errno.RedisGetDataError)
		return
	}

	if len(keys) == 0 {
		log.C(ctx).Info().Msg("redis中不存在以wx_开头的键")
		response.Fail(ctx, errno.RedisNoDataError)
		return
	}

	for _, key := range keys {
		//获取键对应的值
		value, rErr := t.sdk.Rdb().Get(ctx, key).Result()

		if rErr != nil {
			log.C(ctx).Error().Err(err).Str("key", key).Msg("获取redis数据失败")
			continue
		}

		result.Data = append(
			result.Data, v1.TransferBackupItem{
				Key:   key,
				Value: value,
			},
		)
	}

	bByte, err := json.Marshal(result)
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("json序列化失败")
		response.Fail(ctx, errno.JsonEncodeError)
		return
	}

	encrypted, err := help.AesEncrypt(string(bByte), []byte(constant.AesBackup))

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("加密失败")
		response.Fail(ctx, errno.AesEncryptError)
		return
	}

	response.Success(
		ctx, gin.H{
			"backup": encrypted,
		},
	)
}

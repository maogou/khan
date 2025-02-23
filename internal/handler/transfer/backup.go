package transfer

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (t *TransferHandler) Backup(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用TransferHandler->Backup方法")

	var result = v1.TransferResponse{
		Data: make([]v1.TransferBackupItem, 0),
		Tip:  "如果data数据不为空的话，请list对应的数据保存为backup.json,防止后面迁移时候使用，这样可以保证迁移后继续使用之前的appid",
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

	response.Success(ctx, result)
}

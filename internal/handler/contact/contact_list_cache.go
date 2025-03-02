package contact

import (
	"encoding/json"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *ContactHandler) FetchContactsListCache(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ContactHandler->FetchContactsListCache方法")

	var (
		req    v1.ContactListRequest
		result = v1.ContactListResponse{
			Chatrooms: make([]string, 0),
			Friends:   make([]string, 0),
			Ghs:       make([]string, 0),
		}
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	cKey := "contact_cache_" + req.AppId

	contacts, err := c.sdk.Rdb().Get(ctx, cKey).Result()

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("获取联系人列表缓存失败")
		response.Fail(ctx, errno.GetContactListCacheError)
		return
	}

	if err = json.Unmarshal([]byte(contacts), &result); err != nil {
		log.C(ctx).Error().Err(err).Msg("json反序列化失败")
		response.Fail(ctx, errno.JsonDecodeError)
		return
	}

	response.Success(ctx, result)
}

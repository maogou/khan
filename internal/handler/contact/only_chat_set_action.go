package contact

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	v1 "smallBot/api/gewe/v1"
	"smallBot/api/gewe/v1/transform"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
)

func (c *ContactHandler) OnlyChat(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ContactHandler->OnlyChat方法")

	var req v1.ContactSetOnlyChatRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ContactSetOnlyChat(
		ctx, transform.ContactSetOnlyChatRequest{
			Appid: c.sdk.GetAppId(),
			Config: transform.ContactSetOnlyChatConfig{
				Status: lo.Ternary(req.OnlyChat, 1, 2),
				ToWxid: req.Wxid,
			},
			FuncId: 72,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用sdk->ContactSetOnlyChat方法失败")
		response.Fail(ctx, errno.ContactSetOnlyChatError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用sdk->ContactSetOnlyChat方法失败")
		response.Fail(ctx, errno.ContactSetOnlyChatError)
		return
	}

	response.SuccessMsg(ctx, "设置好友权限成功")
}

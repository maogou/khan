package contact

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/contact"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
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
		ctx, contact.ContactSetOnlyChatRequest{
			Appid: req.AppId,
			Config: contact.ContactSetOnlyChatConfig{
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

package contact

import (
	"github.com/gin-gonic/gin"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
)

func (c *ContactHandler) Delete(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ContactHandler->Delete方法")

	var req v1.ContactDelRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ContactDel(
		ctx, transform.ContactDelRequest{
			Appid:  req.AppId,
			ToWxid: req.Wxid,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ContactDel方法失败")
		response.Fail(ctx, errno.ContactDelError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret !=0 ->调用ContactDel方法失败")
		response.Fail(ctx, errno.ContactDelError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")
}

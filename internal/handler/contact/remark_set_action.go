package contact

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/contact"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *ContactHandler) Remark(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ContactHandler->Remark方法")

	var req v1.ContactSetRemarkResquest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ContactSetRemark(
		ctx, contact.ContactSetRemarkRequest{
			Appid:  req.AppId,
			Remark: req.Remark,
			ToWxid: req.Wxid,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ContactSetRemark方法失败")
		response.Fail(ctx, errno.ContactSetRemarkError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret !=0 ->调用ContactSetRemark方法失败")
		response.Fail(ctx, errno.ContactSetRemarkError)
		return
	}

	response.SuccessMsg(ctx, "设置好友备注成功")
}

package contact

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/contact"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c ContactHandler) Add(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ContactHandler->Add方法")

	var req v1.ContactAddRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ContactAdd(
		ctx, contact.ContactAddRequest{
			Appid:   req.AppId,
			Content: req.Content,
			Opcode:  req.Option,
			Scene:   req.Scene,
			V3:      req.V3,
			V4:      req.V4,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ContactAdd方法失败")
		response.Fail(ctx, errno.ContactAddError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret !=0 ->调用ContactAdd方法失败")
		response.Fail(ctx, errno.ContactAddError)
		return
	}

	response.SuccessMsg(ctx, "操作成功，等待对方同意")
}

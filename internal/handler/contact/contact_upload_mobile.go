package contact

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/contact"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *ContactHandler) ContactUploadMobile(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用ContactHandler->ContactUploadMobile方法")

	var req v1.ContactUploadMobileRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := c.sdk.ContactUploadMobile(
		ctx, contact.ContactUploadMobileRequest{
			Appid:   req.AppId,
			Opcode:  req.OpType,
			PhoneNo: req.Phones,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用sdk->ContactUploadMobile方法失败")
		response.Fail(ctx, errno.ContactUploadMobileError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("调用sdk->ContactUploadMobile方法resp.Ret != 0失败")
		response.Fail(ctx, errno.ContactUploadMobileError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Msg("调用sdk->ContactUploadMobile方法resp.Data.BaseResponse.Ret != 0失败")
		response.Fail(ctx, errno.ContactUploadMobileError)
		return
	}

	response.SuccessMsg(ctx, "上传成功")
}

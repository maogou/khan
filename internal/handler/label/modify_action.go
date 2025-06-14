package label

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/label"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (l *LabelHandler) Modify(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用LabelHandler->Modify方法")

	var req v1.LabelModifyRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := l.sdk.LabelModify(
		ctx, label.LabelModifyRequest{
			Appid:      req.AppId,
			Id:         req.LabelIds,
			ToWxidList: req.WxIds,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用LabelModify方法失败")
		response.Fail(ctx, errno.LabelModifyError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用LabelModify方法失败")
		response.Fail(ctx, errno.LabelModifyError)
		return
	}

	response.SuccessMsg(ctx, "修改标签成功")
}

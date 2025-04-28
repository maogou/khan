package label

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/label"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (l *LabelHandler) Delete(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用LabelHandler->Delete方法")

	var req v1.LabelDeleteRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := l.sdk.LabelDelete(
		ctx, label.LabelDeleteRequest{
			Appid: req.AppId,
			Id:    req.LabelIds,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用LabelDelete方法失败")
		response.Fail(ctx, errno.LabelDeleteError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用LabelDelete方法失败")
		response.Fail(ctx, errno.LabelDeleteError)
		return
	}

	response.SuccessMsg(ctx, "删除成功")
}

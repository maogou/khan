package label

import (
	v1 "smallBot/api/gewe/v1"
	"smallBot/api/gewe/v1/transform"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

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
		ctx, transform.LabelDeleteRequest{
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

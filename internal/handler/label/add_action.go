package label

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/label"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (l *LabelHandler) Add(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用LabelHandler->Add方法")

	var req v1.LabelAddRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := l.sdk.LabelAdd(
		ctx, label.LabelAddRequest{
			Appid: req.AppId,
			Name:  req.LabelName,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用LabelAdd方法失败")
		response.Fail(ctx, errno.LabelAddError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用LabelAdd方法失败")
		response.Fail(ctx, errno.LabelAddError)
		return
	}

	response.Success(
		ctx, v1.LabelAddResponse{
			LabelId:   resp.Data.LabelPairList.LabelID,
			LabelName: resp.Data.LabelPairList.LabelName,
		},
	)

}

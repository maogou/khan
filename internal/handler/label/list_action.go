package label

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/label"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (l *LabelHandler) List(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用LabelHandler->List方法")

	var (
		req   v1.LabelListRequest
		items = make([]v1.LabelListItem, 0)
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := l.sdk.LabelList(
		ctx, label.LabelListRequest{
			Appid: req.AppId,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用LabelList方法失败")
		response.Fail(ctx, errno.LabelListError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用LabelList方法失败")
		response.Fail(ctx, errno.LabelListError)
		return
	}

	for _, v := range resp.Data.LabelPairList {
		items = append(
			items, v1.LabelListItem{
				LabelId:   v.LabelID,
				LabelName: v.LabelName,
			},
		)
	}

	response.Success(
		ctx, v1.LabelListResponse{
			LabelList: items,
		},
	)

}

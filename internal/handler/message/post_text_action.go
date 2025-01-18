package message

import (
	v1 "smallBot/api/gewe/v1"
	"smallBot/api/gewe/v1/transform"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (m *MessageHandler) PostText(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->PostText方法")

	var req v1.PostTextRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, err)
		return
	}

	resp, err := m.sdk.PostText(ctx, transform.PostTextRequest{
		AppId: req.AppId,
		ToWxidList: []transform.ToWxid{
			{
				Content: req.Content,
				ToWxid:  req.ToWxid,
				MsgType: 1,
			},
		},
	},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用PostText方法发送消息失败")
		response.Fail(ctx, errno.SendMsgError)
		return
	}

	log.C(ctx).Info().Any("req", req).Any("resp", resp).Msg("发送消息成功")

	response.Success(ctx, v1.PostTextResponse{
		ToWxid:     req.ToWxid,
		MsgId:      resp.Data.List[0].MsgId,
		NewMsgId:   resp.Data.List[0].NewMsgId,
		Type:       resp.Data.List[0].Type,
		CreateTime: resp.Data.List[0].Createtime,
	})

	return
}

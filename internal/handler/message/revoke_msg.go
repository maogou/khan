package message

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/message"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (m *MessageHandler) RevokeMsg(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->RevokeMsg方法")

	var req v1.RevokeMsgRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := m.sdk.RevokeMsg(
		ctx, message.RevokeMsgRequest{
			Appid:      req.AppId,
			CreateTime: req.CreateTime,
			NewMsgId:   req.NewMsgId,
			ToWxid:     req.ToWxid,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用RevokeMsg方法失败")
		response.Fail(ctx, errno.RevokeMsgError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("ret !=0 ->调用RevokeMsg方法失败")
		response.Fail(ctx, errno.RevokeMsgError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")

}

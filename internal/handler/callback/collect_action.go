package callback

import (
	"encoding/json"
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (c *CallbackHandler) Collect(ctx *gin.Context) {
	log.C(ctx).Info().Msg("开始调用CallbackHandler-Collect")
	var (
		req     v1.CollectRequest
		msgType v1.CallbackMessageType
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("CallbackHandler-Collect参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	log.C(ctx).Info().Any("req", req).Msg("接收到原始的回调数据")

	if err := json.Unmarshal(req.Data, &msgType); err != nil {
		log.C(ctx).Error().Err(err).Msg("CallbackHandler-Collect解析msgType失败")
		response.Fail(ctx, errno.CallbackMsgTypeErr)
		return
	}

	req.MsgType = msgType.MsgType

	//中间件种和handler种必须使用ctx的副本
	cCp := ctx.Copy()
	go func() {
		c.chain.HandlerRequest(cCp, req)
	}()

	response.SuccessMsg(ctx, "success")
}

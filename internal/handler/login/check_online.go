package login

import (
	v1 "smallBot/api/khan/v1"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (l *LoginHandler) CheckOnline(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用LoginHandler->CheckOnline方法")

	var req v1.CheckOnlineRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	hResp, err := l.sdk.HearBeat(
		ctx, v1.HearBeatRequest{
			AppId: req.AppId,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用HearBeat方法失败")
		response.Fail(ctx, errno.HearBeatError)
		return
	}

	if hResp.Ret != 0 {
		log.C(ctx).Error().Msg("ret != 0 ->调用HearBeat方法失败")
		response.Fail(ctx, errno.HearBeatError)
		return
	}

	if hResp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Msg("BaseResponse.Ret !=0 ->调用HearBeat方法失败")
		response.Fail(ctx, errno.HearBeatError)
		return
	}

	response.SuccessMsg(ctx, "ok")
}

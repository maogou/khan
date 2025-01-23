package personal

import (
	v1 "smallBot/api/gewe/v1"
	"smallBot/api/gewe/v1/transform"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (p PersonalHandler) Safety(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用PersonalHandler->Safety方法")

	var req v1.PersonalSafetyRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := p.sdk.PersonalSafety(
		ctx, transform.PersonalSafetyRequest{
			AppId:   req.AppId,
			ProxyIp: req.ProxyIp,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用PersonalSafety方法失败")
		response.Fail(ctx, errno.PersonalSafetyError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用PersonalSafety方法失败")
		response.Fail(ctx, errno.PersonalSafetyError)
		return
	}

	var result v1.PersonalSafetResponse

	for _, v := range resp.Data.Info.Devicelist {
		result.List = append(
			result.List, v1.PersonalSafetyItem{
				Uuid:       v.Uuid,
				DeviceName: v.Devicename,
				DeviceType: v.Devicetype,
				LastTime:   v.Lasttime,
			},
		)
	}

	response.Success(ctx, result)
}

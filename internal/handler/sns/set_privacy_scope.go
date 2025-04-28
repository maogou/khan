package sns

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/sns"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (s *SnsHandler) SetPrivacyScope(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->SetPrivacyScope方法")

	var (
		req       v1.SnsSetPrivacyScopeRequest
		optionMap = map[int]int{
			1: 4294967295,
			2: 4320,
			3: 720,
			4: 72,
		}
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := s.sdk.SnsSetPrivacyScope(
		ctx, sns.SnsSetPrivacyScopeRequest{
			AppId: req.AppId,
			Id:    6273,
			Value: optionMap[req.Option],
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsSetPrivacyScope方法失败")
		response.Fail(ctx, errno.SnsSetPrivacyScopeError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用SnsSetPrivacyScope方法失败")
		response.Fail(ctx, errno.SnsSetPrivacyScopeError)
		return
	}

	response.SuccessMsg(ctx, "操作成功")

}

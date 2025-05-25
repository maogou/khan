package sns

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (s *SnsHandler) Detail(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->Detail方法")

	var (
		req v1.SnsDetailRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := s.sdk.SnsDetail(ctx, req)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsDetail方法失败")
		response.Fail(ctx, errno.SnsDetailError)
		return
	}

	if resp.Ret != 200 {
		log.C(ctx).Error().Err(err).Msg("ret != 0 ->调用SnsDetail方法失败")
		response.Fail(ctx, errno.SnsDetailError)
		return
	}

	response.Success(ctx, resp.Data)

}

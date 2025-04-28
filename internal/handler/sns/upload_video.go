package sns

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/sns"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (s *SnsHandler) UploadVideo(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->UploadVideo方法")

	var req v1.SnsUploadVideoRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	resp, err := s.sdk.SnsUploadVideo(
		ctx, sns.SnsUploadVideoRequest{
			AppId:    req.AppId,
			File:     req.VideoUrl,
			Pic:      req.ThumbUrl,
			IsBase64: false,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsUploadVideo方法失败")
		response.Fail(ctx, errno.SnsUploadVideoError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("ret != 0 ->调用SnsUploadVideo方法失败")
		response.Fail(ctx, errno.SnsUploadVideoError)
		return
	}

	response.Success(
		ctx, v1.SnsUploadVideoResponse{
			FileMd5:  resp.Data.FileMd5,
			FileUrl:  resp.Data.FileUrl,
			ThumbUrl: resp.Data.Thumburl,
			Length:   resp.Data.RecvLen,
		},
	)

}

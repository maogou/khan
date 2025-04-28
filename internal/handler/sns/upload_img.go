package sns

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/sns"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (s *SnsHandler) UploadImg(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->UploadImg方法")

	var (
		req    v1.SnsUploadImgRequest
		result = v1.SnsUploadImgResponse{
			Data: []v1.SnsUploadImgItem{},
		}
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	for index, imgUrl := range req.ImgUrls {
		resp, err := s.sdk.SnsUploadImg(
			ctx, sns.SnsUploadImgRequest{
				Appid: req.AppId,
				Config: sns.SnsUploadImgConfig{
					Thumbnail: true,
				},
				FileLink: imgUrl,
			},
		)

		if err != nil {
			log.C(ctx).Error().Err(err).Int("index", index).Msg("调用SnsUploadImg方法失败")
			continue
		}

		if resp.Ret != 0 {
			log.C(ctx).Error().Err(err).Int("index", index).Msg("调用SnsUploadImg方法失败")
			continue
		}

		result.Data = append(
			result.Data, v1.SnsUploadImgItem{
				FileUrl:  resp.Data.FileUrl,
				ThumbUrl: resp.Data.Thumburl,
				FileMd5:  resp.Data.FileMd5,
				Length:   resp.Data.RecvLen,
				Width:    resp.Data.ImgInfo.Width,
				Height:   resp.Data.ImgInfo.Height,
			},
		)
	}

	response.Success(ctx, result)

}

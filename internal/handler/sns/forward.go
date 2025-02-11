package sns

import (
	"encoding/xml"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/sns"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (s *SnsHandler) Forward(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->Forward方法")

	var (
		req            v1.SnsForwardRequest
		transformImage []sns.SnsSendImageItem
		timeline       sns.ImageTimelineXml
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	if err := xml.Unmarshal([]byte(req.SnsXml), &timeline); err != nil {
		log.C(ctx).Error().Err(err).Msg("xml反序列化失败")
		response.Fail(ctx, errno.XmlDecodeError)
	}

	if timeline.ContentObject.ContentStyle == 1 {
		log.C(ctx).Info().Any("timeline", timeline).Msg("timeline")

		for _, media := range timeline.ContentObject.MediaList.Media {
			transformImage = append(
				transformImage, sns.SnsSendImageItem{
					ThumbUrl:  media.Thumb.Text,
					Width:     media.Size.Width,
					TotalSize: media.Size.TotalSize,
					Id:        media.ID,
					Type:      media.Type,
					Url:       media.URL.Text,
					Height:    media.Size.Height,
					Md5:       media.URL.MD5,
				},
			)
		}

		response.Success(ctx, transformImage)
	}

	log.C(ctx).Info().Any("timeline", timeline).Msg("timeline")

	response.SuccessMsg(ctx, req.SnsXml)
}

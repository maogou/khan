package sns

import (
	"encoding/xml"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/sns"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
	"strconv"

	"github.com/samber/lo"
	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
)

func (s *SnsHandler) Forward(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->Forward方法")

	var (
		req      v1.SnsForwardRequest
		timeline v1.TimelineObject
		result   *v1.SnsForwardResponse
		err      error
	)

	if err = ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	if err = xml.Unmarshal([]byte(req.SnsXml), &timeline); err != nil {
		log.C(ctx).Error().Err(err).Msg("xml反序列化失败")
		response.Fail(ctx, errno.XmlDecodeError)
		return
	}

	contentStyle := timeline.ContentObject.ContentStyle

	switch contentStyle {
	case 1:

		result, err = s.snsSendImage(ctx, req, timeline)
	default:

		result, err = s.snsSendExcludeImage(ctx, req)
	}

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsHandler->Forward方法失败")
		response.Fail(ctx, errno.SnsForwardError)
		return
	}

	response.Success(ctx, result)
}

func (s *SnsHandler) snsSendImage(ctx *gin.Context, req v1.SnsForwardRequest, timeline v1.TimelineObject) (*v1.SnsForwardResponse, error) {
	log.C(ctx).Info().Msg("调用SnsHandler->snsSendImage方法")

	var transformImage []sns.SnsSendImageItem

	for _, media := range timeline.ContentObject.MediaList.Media {
		transformImage = append(
			transformImage, sns.SnsSendImageItem{
				ThumbUrl:  media.Thumb.Text,
				Width:     strconv.Itoa(cast.ToInt(media.Size.Width)),
				TotalSize: media.Size.TotalSize,
				Id:        media.Id,
				Type:      "1",
				Url:       media.Url.Text,
				Height:    strconv.Itoa(cast.ToInt(media.Size.Height)),
				Md5:       media.Url.Md5,
			},
		)
	}

	resp, err := s.sdk.SnsSendImage(
		ctx, sns.SnsSendImageRequest{
			AppId:        req.AppId,
			AllowUser:    lo.Ternary(len(req.AllowWxIds) > 0, req.AllowWxIds, make([]string, 0)),
			AllowTagId:   make([]string, 0),
			AtUser:       lo.Ternary(len(req.AtWxIds) > 0, req.AtWxIds, make([]string, 0)),
			DisableUser:  lo.Ternary(len(req.DisableWxIds) > 0, req.DisableWxIds, make([]string, 0)),
			DisableTagId: make([]string, 0),
			Private:      req.Privacy,
			XmlTxt:       req.SnsXml,
			Content:      timeline.ContentDesc,
			Media:        transformImage,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsHandler->snsSendImage方法失败")
		return nil, errno.SnsSendImageError
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("ret !=0 ->调用SnsHandler->snsSendImage方法失败")
		return nil, errno.SnsSendImageError
	}

	return &v1.SnsForwardResponse{
		Id:         resp.Data.SnsObject.Id,
		UserName:   resp.Data.SnsObject.Username,
		NickName:   resp.Data.SnsObject.Nickname,
		CreateTime: resp.Data.SnsObject.CreateTime,
	}, err
}

func (s *SnsHandler) snsSendExcludeImage(ctx *gin.Context, req v1.SnsForwardRequest) (*v1.SnsForwardResponse, error) {
	log.C(ctx).Info().Msg("调用SnsHandler->snsSendExcludeImage方法")

	resp, err := s.sdk.SnsSendExcludeImage(
		ctx, sns.SnsSendExcludeImageRequest{
			AppId:        req.AppId,
			AllowUser:    lo.Ternary(len(req.AllowWxIds) > 0, req.AllowWxIds, make([]string, 0)),
			AllowTagId:   make([]string, 0),
			AtUser:       lo.Ternary(len(req.AtWxIds) > 0, req.AtWxIds, make([]string, 0)),
			DisableUser:  lo.Ternary(len(req.DisableWxIds) > 0, req.DisableWxIds, make([]string, 0)),
			DisableTagId: make([]string, 0),
			Private:      req.Privacy,
			XmlTxt:       req.SnsXml,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsHandler->snsSendExcludeImage方法失败")
		return nil, errno.SnsSendExcludeImageError
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("ret !=0 ->调用SnsHandler->snsSendExcludeImage方法失败")
		return nil, errno.SnsSendExcludeImageError
	}

	return &v1.SnsForwardResponse{
		Id:         resp.Data.SnsObject.Id,
		UserName:   resp.Data.SnsObject.Username,
		NickName:   resp.Data.SnsObject.Nickname,
		CreateTime: resp.Data.SnsObject.CreateTime,
	}, err

}

package sns

import (
	"encoding/xml"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/sns"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *SnsHandler) Forward(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->Forward方法")

	var (
		req            v1.SnsForwardRequest
		transformImage []sns.SnsSendImageItem
		timeline       v1.TimelineObject
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	if err := xml.Unmarshal([]byte(req.SnsXml), &timeline); err != nil {
		log.C(ctx).Error().Err(err).Msg("xml反序列化失败")
		response.Fail(ctx, errno.XmlDecodeError)
		return
	}

	if timeline.ContentObject.ContentStyle == 1 {
		log.C(ctx).Info().Any("timeline", timeline).Msg("转发的是图片类型")

		for _, media := range timeline.ContentObject.MediaList.Media {
			transformImage = append(
				transformImage, sns.SnsSendImageItem{
					ThumbUrl:  media.Thumb.Text,
					Width:     strconv.Itoa(cast.ToInt(media.Size.Width)),
					TotalSize: media.Size.TotalSize,
					Id:        media.Id,
					Type:      media.Type,
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
			log.C(ctx).Error().Err(err).Msg("调用SnsHandler->ForwardImage方法失败")
			response.Fail(ctx, errno.SnsSendImageError)
			return
		}

		if resp.Ret != 0 {
			log.C(ctx).Error().Msg("ret !=0 ->调用SnsHandler->ForwardImage方法失败")
			response.Fail(ctx, errno.SnsSendImageError)
			return
		}

		if resp.Data.BaseResponse.Ret != 0 {
			log.C(ctx).Error().Any(
				"msg_err", resp.Data.BaseResponse.ErrMsg,
			).Msg("BaseResponse.Ret !=0 ->调用SnsHandler->ForwardImage方法失败")
			response.Fail(ctx, errno.SnsSendImageError)
			return
		}

		response.Success(
			ctx, v1.SnsSendImageResponse{
				Id:         resp.Data.SnsObject.Id,
				UserName:   resp.Data.SnsObject.Username,
				NickName:   resp.Data.SnsObject.Nickname,
				CreateTime: resp.Data.SnsObject.CreateTime,
			},
		)

		return
	}

	resp, err := s.sdk.SnsSendText(
		ctx, sns.SnsSendTextRequest{
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
		log.C(ctx).Error().Err(err).Msg("调用SnsHandler->ForwardText方法失败")
		response.Fail(ctx, errno.SnsSendTextError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("ret !=0 ->调用SnsHandler->ForwardText方法失败")
		response.Fail(ctx, errno.SnsSendTextError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Any(
			"msg_err", resp.Data.BaseResponse.ErrMsg,
		).Msg("BaseResponse.Ret !=0 ->调用SnsHandler->ForwardText方法失败")
		response.Fail(ctx, errno.SnsSendTextError)
		return
	}

	response.Success(
		ctx, v1.SnsSendTextResponse{
			Id:         resp.Data.SnsObject.Id,
			UserName:   resp.Data.SnsObject.Username,
			NickName:   resp.Data.SnsObject.Nickname,
			CreateTime: resp.Data.SnsObject.CreateTime,
		},
	)
}

package sns

import (
	"bytes"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/sns"
	"maogou/khan/internal/constant"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func (s *SnsHandler) SendVideo(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->SendVideo方法")

	var (
		req           v1.SnsSendVideoRequest
		videoTimeline sns.SnsVideoTimelineObject
		buf           bytes.Buffer
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	videoTimeline.Content = req.Content
	videoTimeline.Video.FileUrl = req.VideoInfo.FileUrl
	videoTimeline.Video.VideoDuration = strconv.Itoa(req.VideoInfo.Length)
	videoTimeline.Video.ThumbUrl = req.VideoInfo.ThumbUrl

	t := template.Must(template.New("video_timeline").Parse(constant.SnsVideo))

	if err := t.Execute(&buf, videoTimeline); err != nil {
		log.C(ctx).Error().Err(err).Msg("模板解析失败")
		response.Fail(ctx, errno.TemplateParseError)
		return
	}

	resp, err := s.sdk.SnsSendVideo(
		ctx, sns.SnsSendVideoRequest{
			AppId:        req.AppId,
			AllowUser:    lo.Ternary(len(req.AllowWxIds) > 0, req.AllowWxIds, make([]string, 0)),
			AllowTagId:   make([]string, 0),
			AtUser:       lo.Ternary(len(req.AtWxIds) > 0, req.AtWxIds, make([]string, 0)),
			DisableUser:  lo.Ternary(len(req.DisableWxIds) > 0, req.DisableWxIds, make([]string, 0)),
			DisableTagId: make([]string, 0),
			Private:      req.Privacy,
			XmlTxt:       buf.String(),
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsSendVideo方法失败")
		response.Fail(ctx, errno.SnsSendVideoError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("ret !=0 ->调用SnsSendVideo方法失败")
		response.Fail(ctx, errno.SnsSendVideoError)
		return
	}

	response.Success(
		ctx, v1.SnsSendVideoResponse{
			Id:         resp.Data.SnsObject.Id,
			UserName:   resp.Data.SnsObject.Username,
			NickName:   resp.Data.SnsObject.Nickname,
			CreateTime: resp.Data.SnsObject.CreateTime,
		},
	)
}

package sns

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/sns"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
	"text/template"
)

func (s *SnsHandler) SendUrl(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->SendUrl方法")

	var (
		req         v1.SnsSendUrlRequest
		urlTimeline sns.UrlTimelineObject
		buf         bytes.Buffer
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	urlTimeline.Content = req.Content
	urlTimeline.Title = req.Title
	urlTimeline.ContentUrl = req.LinkUrl
	urlTimeline.Description = req.Description
	urlTimeline.ThumbUrl = req.ThumbUrl

	t := template.Must(template.New("url_timeline").Parse(constant.SnsUrl))

	if err := t.Execute(&buf, urlTimeline); err != nil {
		log.C(ctx).Error().Err(err).Msg("模板解析失败")
		response.Fail(ctx, errno.TemplateParseError)
		return
	}

	resp, err := s.sdk.SnsSendUrl(
		ctx, sns.SnsSendUrlRequest{
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
		log.C(ctx).Error().Err(err).Msg("调用SnsHandler->SendUrl方法失败")
		response.Fail(ctx, errno.SnsSendUrlError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("ret !=0 ->调用SnsHandler->SendUrl方法失败")
		response.Fail(ctx, errno.SnsSendUrlError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Any(
			"msg_err", resp.Data.BaseResponse.ErrMsg,
		).Msg("BaseResponse.Ret !=0 ->调用SnsHandler->SendUrl方法失败")
		response.Fail(ctx, errno.SnsSendUrlError)
		return
	}

	response.Success(
		ctx, v1.SnsSendUrlResponse{
			Id:         resp.Data.SnsObject.Id,
			UserName:   resp.Data.SnsObject.Username,
			NickName:   resp.Data.SnsObject.Nickname,
			CreateTime: resp.Data.SnsObject.CreateTime,
		},
	)
}

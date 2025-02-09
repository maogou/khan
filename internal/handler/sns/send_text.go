package sns

import (
	"encoding/xml"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/sns"
	"smallBot/api/khan/v1/transform/sns/template"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/samber/lo"

	"github.com/gin-gonic/gin"
)

func (s *SnsHandler) SendText(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SnsHandler->SendText方法")

	var (
		req         v1.SnsSendTextRequest
		xmlTemplate template.TimelineObject
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	xmlTemplate.ContentDesc = template.CDATA{
		Content: req.Content,
	}

	xmlByte, err := xml.Marshal(xmlTemplate)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("xml序列化失败")
		response.Fail(ctx, errno.XmlEncodeError)
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
			XmlTxt:       string(xmlByte),
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SnsSendText方法失败")
		response.Fail(ctx, errno.SnsSendTextError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("ret !=0 ->调用SnsSendText方法失败")
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

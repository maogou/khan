package message

import (
	"bytes"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/message"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
	"text/template"

	"github.com/gin-gonic/gin"
)

func (m *MessageHandler) SendMiniApp(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->SendMiniApp方法")

	var (
		req         v1.SendMiniAppRequest
		sendMiniTpl message.SendMiniApp
		buf         bytes.Buffer
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	sendMiniTpl.Title = req.Title
	sendMiniTpl.AppId = req.AppId
	sendMiniTpl.DisplayName = req.DisplayName
	sendMiniTpl.PagePath = req.PagePath
	sendMiniTpl.UserName = req.ToWxid

	t := template.Must(template.New("send_mini_app").Parse(constant.SendApplet))

	if err := t.Execute(&buf, sendMiniTpl); err != nil {
		log.C(ctx).Error().Err(err).Msg("模板解析失败")
		response.Fail(ctx, errno.TemplateParseError)
		return
	}

	resp, err := m.sdk.SendMiniApp(
		ctx, message.SendMiniAppRequest{
			Appid:    req.AppId,
			ToWxid:   req.ToWxid,
			Xml:      buf.String(),
			FileLink: req.CoverImgUrl,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用SendMiniApp方法发送消息失败")
		response.Fail(ctx, errno.SendMiniAppError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用SendMiniApp方法发送消息失败")
		response.Fail(ctx, errno.SendMiniAppError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用SendMiniApp-resp.Data.BaseResponse.Ret != 0方法发送消息失败")
		response.Fail(ctx, errno.SendMiniAppError)
		return
	}

	response.Success(
		ctx, v1.SendMiniAppResponse{
			ToWxid:     req.ToWxid,
			MsgId:      resp.Data.MsgId,
			Type:       resp.Data.Type,
			NewMsgId:   resp.Data.NewMsgId,
			CreateTime: resp.Data.CreateTime,
		},
	)
}

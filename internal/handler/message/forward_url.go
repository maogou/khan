package message

import (
	"encoding/xml"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/message"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

type ForwardUrlMsg struct {
	XMLName xml.Name `xml:"appmsg"`
	Content string   `xml:",innerxml"` // 获取 appmsg 节点下的所有内容
}

type UrlRoot struct {
	XMLName xml.Name          `xml:"msg"`
	AppMsg  ForwardMiniAppMsg `xml:"appmsg"`
}

func (m *MessageHandler) ForwardUrl(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->ForwardUrl方法")

	var (
		req     v1.ForwardUrlRequest
		urlRoot UrlRoot
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	xs := strings.NewReplacer(
		"\n", "",
		"\t", "",
		`\`, "",
	).Replace(req.Xml)

	if err := xml.Unmarshal([]byte(xs), &urlRoot); err != nil {
		log.C(ctx).Error().Err(err).Msg("xml反序列化失败")
		response.Fail(ctx, errno.XmlDecodeError)
		return
	}

	appMsg, err := xml.Marshal(urlRoot.AppMsg)
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("xml序列化失败")
		response.Fail(ctx, errno.XmlEncodeError)
		return
	}

	resp, err := m.sdk.ForwardUrl(
		ctx, message.ForwardUrlRequest{
			Appid:  req.AppId,
			ToWxid: req.ToWxid,
			Xml:    string(appMsg),
			Type:   5,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ForwardUrl方法失败")
		response.Fail(ctx, errno.ForwardUrlError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("调用ForwardUrl方法resp.Ret != 0失败")
		response.Fail(ctx, errno.ForwardUrlError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Msg("调用ForwardUrl方法resp.Data.BaseResponse.Ret != 0失败")
		response.Fail(ctx, errno.ForwardUrlError)
		return
	}

	response.Success(ctx, v1.ForwardUrlResponse{
		CreateTime: resp.Data.CreateTime,
		MsgId:      resp.Data.MsgId,
		NewMsgId:   resp.Data.NewMsgId,
		ToWxid:     resp.Data.ToUserName,
		Type:       resp.Data.Type,
	})

}

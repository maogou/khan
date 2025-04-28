package message

import (
	"encoding/xml"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/message"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

type ForwardMiniAppMsg struct {
	XMLName xml.Name `xml:"appmsg"`
	Content string   `xml:",innerxml"` // 获取 appmsg 节点下的所有内容
}

type Root struct {
	XMLName xml.Name          `xml:"msg"`
	AppMsg  ForwardMiniAppMsg `xml:"appmsg"`
}

func (m *MessageHandler) ForwardMiniApp(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->ForwardMiniApp方法")

	var (
		req  v1.ForwardMiniAppRequest
		root Root
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

	if err := xml.Unmarshal([]byte(xs), &root); err != nil {
		log.C(ctx).Error().Err(err).Msg("xml反序列化失败")
		response.Fail(ctx, errno.XmlDecodeError)
		return
	}

	appMsg, err := xml.Marshal(root.AppMsg)
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("xml序列化失败")
		response.Fail(ctx, errno.XmlEncodeError)
		return
	}

	resp, err := m.sdk.ForwardMiniApp(
		ctx, message.ForwardMiniAppRequest{
			Appid:    req.AppId,
			FileLink: req.CoverImgUrl,
			ToWxid:   req.ToWxid,
			Xml:      string(appMsg),
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ForwardMiniApp方法发送消息失败")
		response.Fail(ctx, errno.ForwardMiniAppError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Any("resp", resp).Msg("调用ForwardMiniApp方法发送消息失败")
		response.Fail(ctx, errno.ForwardMiniAppError)
		return
	}

	response.Success(
		ctx, v1.ForwardMiniAppResponse{
			CreateTime: resp.Data.CreateTime,
			MsgId:      resp.Data.MsgId,
			NewMsgId:   resp.Data.NewMsgId,
			ToWxid:     resp.Data.ToUserName,
			Type:       resp.Data.Type,
		},
	)
}

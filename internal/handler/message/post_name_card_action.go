package message

import (
	"encoding/xml"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/message"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (m *MessageHandler) PostNameCard(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->PostNameCard方法")

	var (
		req  v1.PostNameCardRequest
		cXml message.NameCardMsgXml
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	cXml.Username = req.NameCardWxid
	cXml.Nickname = req.NickName
	cXml.Alias = req.NameCardWxid
	cXml.ImageStatus = 3
	cXml.Scene = 17
	cXml.Sex = 1
	cXml.RegionCode = "CN"

	xmlByte, err := xml.Marshal(cXml)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("xml序列化失败")
		response.Fail(ctx, errno.XmlEncodeError)
		return
	}

	resp, err := m.sdk.PostNameCard(
		ctx, message.PostNameCardRequest{
			Appid: req.AppId,
			ToWxidList: []message.PostNameCardConfig{
				{
					Content: string(xmlByte),
					ToWxid:  req.NameCardWxid,
					MsgType: 42,
				},
			},
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用PostNameCard方法发送消息失败")
		response.Fail(ctx, errno.PostNameCardError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用PostNameCard方法发送消息失败")
		response.Fail(ctx, errno.PostNameCardError)
		return
	}

	response.Success(
		ctx, v1.PostNameCardResponse{
			ToWxid:     req.ToWxid,
			MsgId:      resp.Data.List[0].MsgId,
			NewMsgId:   resp.Data.List[0].NewMsgId,
			Type:       resp.Data.List[0].Type,
			CreateTime: resp.Data.List[0].Createtime,
		},
	)

}

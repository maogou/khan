package message

import (
	"encoding/xml"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/message"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type Msg struct {
	XMLName  xml.Name `xml:"msg"`
	AppMsg   AppMsg   `xml:"appmsg"`
	FromUser string   `xml:"fromusername"`
}

type AppMsg struct {
	XMLName   xml.Name  `xml:"appmsg"`
	Title     string    `xml:"title"`
	Type      int       `xml:"type"`
	AppAttach AppAttach `xml:"appattach"`
	MD5       string    `xml:"md5"`
}

type AppAttach struct {
	XMLName           xml.Name `xml:"appattach"`
	AttachID          string   `xml:"attachid"`
	CDNAttachURL      string   `xml:"cdnattachurl"`
	TotalLen          int      `xml:"totallen"`
	AESKey            string   `xml:"aeskey"`
	EncryVer          int      `xml:"encryver"`
	FileExt           string   `xml:"fileext"`
	IsLargeFileMsg    int      `xml:"islargefilemsg"`
	OverwriteNewMsgID string   `xml:"overwrite_newmsgid"`
	FileUploadToken   string   `xml:"fileuploadtoken"`
}

func (m *MessageHandler) ForwardFile(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->ForwardFile方法")

	var (
		req    v1.ForwardFileRequest
		xmlMsg Msg
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	if err := xml.Unmarshal([]byte(req.Xml), &xmlMsg); err != nil {
		log.C(ctx).Error().Err(err).Msg("xml反序列化失败")
		response.Fail(ctx, errno.XmlDecodeError)
	}

	resp, err := m.sdk.ForwardFile(
		ctx, message.ForwardFileRequest{
			Appid:        req.AppId,
			ToWxid:       req.ToWxid,
			Cdnattachurl: xmlMsg.AppMsg.AppAttach.CDNAttachURL,
			Md5:          xmlMsg.AppMsg.MD5,
			Title:        xmlMsg.AppMsg.Title,
			Totallen:     xmlMsg.AppMsg.AppAttach.TotalLen,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用ForwardFile方法发送消息失败")
		response.Fail(ctx, errno.ForwardFileError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Any("resp", resp).Msg("调用ForwardFile方法发送消息失败")
		response.Fail(ctx, errno.ForwardFileError)
		return
	}

	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Any(
			"msg_err", resp.Data.BaseResponse.ErrMsg,
		).Msg("BaseResponse.Ret !=0 ->调用ForwardFile方法失败")
		response.Fail(ctx, errno.ForwardFileError)
		return
	}

	response.Success(
		ctx, v1.ForwardFileResponse{
			ToWxid:     req.ToWxid,
			CreateTime: resp.Data.CreateTime,
			MsgId:      resp.Data.MsgId,
			NewMsgId:   resp.Data.NewMsgId,
			Type:       resp.Data.Type,
		},
	)

}

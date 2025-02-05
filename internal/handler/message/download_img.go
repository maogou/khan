package message

import (
	"encoding/xml"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
	"strings"

	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
)

type XmlMsg struct {
	XMLName           xml.Name      `xml:"msg"`
	Img               Img           `xml:"img"`
	PlatformSignature string        `xml:"platform_signature"`
	ImgDataHash       string        `xml:"imgdatahash"`
	ImgSourceInfo     ImgSourceInfo `xml:"ImgSourceInfo"`
}

type Img struct {
	AesKey          string `xml:"aeskey,attr"`
	EncryVer        string `xml:"encryver,attr"`
	CdnThumbAesKey  string `xml:"cdnthumbaeskey,attr"`
	CdnThumbUrl     string `xml:"cdnthumburl,attr"`
	CdnThumbLength  string `xml:"cdnthumblength,attr"`
	CdnThumbHeight  string `xml:"cdnthumbheight,attr"`
	CdnThumbWidth   string `xml:"cdnthumbwidth,attr"`
	CdnMidHeight    string `xml:"cdnmidheight,attr"`
	CdnMidWidth     string `xml:"cdnmidwidth,attr"`
	CdnHdHeight     string `xml:"cdnhdheight,attr"`
	CdnHdWidth      string `xml:"cdnhdwidth,attr"`
	CdnMidImgUrl    string `xml:"cdnmidimgurl,attr"`
	Length          string `xml:"length,attr"`
	Md5             string `xml:"md5,attr"`
	HevcMidSize     string `xml:"hevc_mid_size,attr"`
	OriginSourceMd5 string `xml:"originsourcemd5,attr"`
}

type ImgSourceInfo struct {
	ImgSourceUrl string `xml:"ImgSourceUrl"`
	BizType      string `xml:"BizType"`
}

func (m *MessageHandler) DownloadImg(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用MessageHandler->DownloadImg方法")

	var (
		req    v1.DownloadImgRequest
		xmlImg XmlMsg
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("参数验证失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	xs := strings.ReplaceAll(req.Xml, `\`, "")
	xs = strings.ReplaceAll(xs, "\t", "")
	xs = strings.ReplaceAll(xs, "\n", "")

	if err := xml.Unmarshal([]byte(xs), &xmlImg); err != nil {
		log.C(ctx).Error().Err(err).Msg("xml反序列化失败")
		response.Fail(ctx, errno.XmlDecodeError)
		return
	}

	resp, err := m.sdk.DownloadImg(
		ctx, transform.DownloadImgRequest{
			AesKey:    xmlImg.Img.AesKey,
			Appid:     req.AppId,
			FileType:  req.Type,
			FileUrl:   xmlImg.Img.CdnMidImgUrl,
			Totalsize: cast.ToInt(xmlImg.Img.Length),
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用DownloadImg方法下载图片失败")
		response.Fail(ctx, errno.DownloadImgError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Any("resp", resp).Msg("调用DownloadImg方法下载图片失败")
		response.Fail(ctx, errno.DownloadImgError)
		return
	}

	response.SuccessMsg(ctx, resp.Data.FileData)
}

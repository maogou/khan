package message

import (
	"encoding/xml"
	"github.com/rs/xid"
	"github.com/samber/lo"
	"os"
	"path/filepath"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/help"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
	"strings"
	"time"

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

	xs := strings.ReplaceAll(req.Xml, "\n", "")
	xs = strings.ReplaceAll(xs, "\t", "")
	xs = strings.ReplaceAll(xs, `\`, "")

	if err := xml.Unmarshal([]byte(xs), &xmlImg); err != nil {
		log.C(ctx).Error().Err(err).Msg("xml反序列化失败")
		response.Fail(ctx, errno.XmlDecodeError)
		return
	}

	cdnThumbAesKey := xmlImg.Img.CdnThumbAesKey
	xmlAesKey := xmlImg.Img.AesKey
	cdnThumbUrl := xmlImg.Img.CdnThumbUrl
	cdnMidImgUrl := xmlImg.Img.CdnMidImgUrl
	aesKey := lo.Ternary(len(xmlAesKey) == 0, cdnThumbAesKey, xmlAesKey)
	fileId := lo.Ternary(len(cdnThumbUrl) == 0, cdnMidImgUrl, cdnThumbUrl)

	execPath, err := os.Executable()
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("获取程序路径失败")
		response.Fail(ctx, errno.ExecPathError)
		return
	}

	log.C(ctx).Info().Msg("执行路径：" + filepath.Dir(execPath))

	resp, err := m.sdk.DownloadImg(
		ctx, transform.DownloadImgRequest{
			AesKey:    aesKey,
			Appid:     req.AppId,
			FileType:  req.Type,
			FileUrl:   fileId,
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

	filename := xid.New().String() + ".png"
	path := filepath.Dir(execPath) + "/public/download/" + time.Now().Local().Format("20060102") + "/" + req.AppId + "/" + filename

	log.C(ctx).Info().Str("path", path).Msg("图片保存的目录")

	if err = help.DownloadBase64File(resp.Data.FileData, path); err != nil {
		log.C(ctx).Error().Err(err).Msg("下载图片失败")
		response.Fail(ctx, errno.DownloadImgError)
		return
	}

	response.Success(
		ctx, v1.DownloadImgResponse{
			FileUrl: path,
		},
	)
}

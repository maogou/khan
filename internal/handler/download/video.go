package download

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/samber/lo"
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/api/khan/v1/transform/download"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/help"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Msg struct {
	XMLName  xml.Name `xml:"msg"`
	VideoMsg VideoMsg `xml:"videomsg"`
}

type VideoMsg struct {
	AESKey            string `xml:"aeskey,attr"`
	CDNVideoURL       string `xml:"cdnvideourl,attr"`
	CDNThumbAESKey    string `xml:"cdnthumbaeskey,attr"`
	CDNThumbURL       string `xml:"cdnthumburl,attr"`
	Length            int    `xml:"length,attr"`
	PlayLength        int    `xml:"playlength,attr"`
	CDNThumbLength    int    `xml:"cdnthumblength,attr"`
	CDNThumbWidth     int    `xml:"cdnthumbwidth,attr"`
	CDNThumbHeight    int    `xml:"cdnthumbheight,attr"`
	FromUsername      string `xml:"fromusername,attr"`
	MD5               string `xml:"md5,attr"`
	NewMD5            string `xml:"newmd5,attr"`
	IsPlaceholder     int    `xml:"isplaceholder,attr"`
	RawMD5            string `xml:"rawmd5,attr"`
	RawLength         int    `xml:"rawlength,attr"`
	CDNRawVideoURL    string `xml:"cdnrawvideourl,attr"`
	CDNRawVideoAESKey string `xml:"cdnrawvideoaeskey,attr"`
	OverwriteNewMsgID int    `xml:"overwritenewmsgid,attr"`
	OriginSourceMD5   string `xml:"originsourcemd5,attr"`
	IsAd              int    `xml:"isad,attr"`
}

func (d *DownloadHandler) DownloadVideo(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用DownloadHandler->DownloadVideo方法")

	var (
		req v1.DownloadVideoRequest
		msg Msg
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("DownloadHandler->DownloadVideo方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	execPath, err := os.Executable()
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("获取程序路径失败")
		response.Fail(ctx, errno.ExecPathError)
		return
	}

	xc := req.Xml

	xc = strings.NewReplacer(
		"\t", "",
		"\n", "",
	).Replace(xc)

	if err := xml.Unmarshal([]byte(xc), &msg); err != nil {
		log.C(ctx).Error().Err(err).Msg("DownloadHandler->DownloadVideo方法解析xml失败")
		response.Fail(ctx, errno.XmlDecodeError)
		return
	}

	aesKey := msg.VideoMsg.AESKey
	cdnAesKey := msg.VideoMsg.CDNThumbAESKey
	videoUrl := msg.VideoMsg.CDNVideoURL
	cdnVideoUrl := msg.VideoMsg.CDNThumbURL

	resp, err := d.sdk.DownloadVideo(
		ctx, download.DownloadVideoRequest{
			Appid:     req.AppId,
			AesKey:    lo.Ternary(len(aesKey) > 0, aesKey, cdnAesKey),
			FileType:  4,
			FileUrl:   lo.Ternary(len(videoUrl) > 0, videoUrl, cdnVideoUrl),
			TotalSize: msg.VideoMsg.Length,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用sdk->DownloadVideo方法失败")
		response.Fail(ctx, errno.DownloadVideoError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("resp.Ret != 0 调用sdk->DownloadVideo方法失败")
		response.Fail(ctx, errno.DownloadVideoError)
		return
	}

	filename := xid.New().String() + ".mp4"
	path := filepath.Dir(execPath) + "/public/download/" + time.Now().Local().Format("20060102") + "/" + req.AppId + "/" + filename

	log.C(ctx).Info().Str("path", path).Msg("视频保存的目录")

	if err = help.DownloadBase64File(resp.Data.FileData, path); err != nil {
		log.C(ctx).Error().Err(err).Msg("下载视频失败")
		response.Fail(ctx, errno.DownloadImgError)
		return
	}

	parts := strings.Split(path, "public/download")
	fileUrl := "http://localhost:" + strconv.Itoa(d.sdk.Config().Port) + parts[1]

	response.Success(
		ctx, v1.DownloadVideoResponse{
			FileUrl: fileUrl,
		},
	)
}

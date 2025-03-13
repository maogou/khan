package download

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"os"
	"path/filepath"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/download"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/help"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
	"strconv"
	"strings"
	"time"
)

type Voice struct {
	XMLName  xml.Name `xml:"msg"`
	VoiceMsg VoiceMsg `xml:"voicemsg"`
}

type VoiceMsg struct {
	EndFlag      int    `xml:"endflag,attr"`
	CancelFlag   int    `xml:"cancelflag,attr"`
	ForwardFlag  int    `xml:"forwardflag,attr"`
	VoiceFormat  int    `xml:"voiceformat,attr"`
	VoiceLength  int    `xml:"voicelength,attr"`
	Length       int    `xml:"length,attr"`
	BufID        int    `xml:"bufid,attr"`
	AesKey       string `xml:"aeskey,attr"`
	VoiceURL     string `xml:"voiceurl,attr"`
	VoiceMD5     string `xml:"voicemd5,attr"`
	ClientMsgID  string `xml:"clientmsgid,attr"`
	FromUserName string `xml:"fromusername,attr"`
}

func (d *DownloadHandler) DownloadVoice(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用DownloadHandler->DownloadVoice")

	var (
		req v1.DownloadVoiceRequest
		msg Voice
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("DownloadHandler->DownloadVoice参数绑定失败")
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

	if err = xml.Unmarshal([]byte(xc), &msg); err != nil {
		log.C(ctx).Error().Err(err).Msg("DownloadHandler->DownloadVoice方法解析xml失败")
		response.Fail(ctx, errno.XmlDecodeError)
		return
	}

	resp, err := d.sdk.DownloadVoice(
		ctx, download.DownloadVoiceRequest{
			Appid:        req.AppId,
			MsgId:        strconv.Itoa(req.MsgId),
			Bufid:        strconv.Itoa(msg.VoiceMsg.BufID),
			Length:       msg.VoiceMsg.Length,
			FromUserName: msg.VoiceMsg.FromUserName,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用DownloadHandler->DownloadVoice方法失败")
		response.Fail(ctx, errno.DownloadVoiceError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Msg("ret !=0 ->DownloadHandler->DownloadVoice方法调用失败")
		response.Fail(ctx, errno.DownloadVoiceError)
		return
	}
	if resp.Data.BaseResponse.Ret != 0 {
		log.C(ctx).Error().Any(
			"msg_err", resp.Data.BaseResponse.ErrMsg,
		).Msg("BaseResponse.Ret !=0 ->调用DownloadVoice方法失败")
		response.Fail(ctx, errno.DownloadVoiceError)
		return
	}

	if len(resp.Data.Data.Buffer) == 0 {
		log.C(ctx).Error().Msg("DownloadHandler->DownloadVoice方法返回的Buffer为空")
		response.Fail(ctx, errno.DownloadVoiceError)
		return
	}

	filename := xid.New().String() + ".silk"
	path := filepath.Dir(execPath) + "/public/download/" + time.Now().Local().Format("20060102") + "/" + req.AppId + "/" + filename

	log.C(ctx).Info().Str("path", path).Msg("语音保存的目录")

	if err = help.DownloadBase64File(resp.Data.Data.Buffer, path); err != nil {
		log.C(ctx).Error().Err(err).Msg("下载语音失败")
		response.Fail(ctx, errno.DownloadImgError)
		return
	}

	parts := strings.Split(path, "public/download")
	fileUrl := "http://localhost:" + strconv.Itoa(d.sdk.Config().Port) + parts[1]

	response.Success(
		ctx, v1.DownloadVoiceResponse{
			FileUrl: fileUrl,
		},
	)
}

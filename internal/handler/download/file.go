package download

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	v1 "smallBot/api/khan/v1"
	"smallBot/api/khan/v1/transform/download"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/help"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"
	"strings"
	"time"
)

type FileMsg struct {
	XMLName xml.Name `xml:"msg"`    // XML根元素名
	AppMsg  AppMsg   `xml:"appmsg"` // 嵌套的应用消息结构
}

type AppMsg struct {
	AppID  string    `xml:"appid,attr"`  // 应用ID，属性形式
	SDKVer string    `xml:"sdkver,attr"` // SDK版本，属性形式
	Title  string    `xml:"title"`       // 标题文本节点
	Attach AppAttach `xml:"appattach"`   // 附件信息嵌套结构
}

type AppAttach struct {
	AttachID          string `xml:"attachid"`           // 附件唯一标识
	CDNAttachURL      string `xml:"cdnattachurl"`       // CDN资源地址
	TotalLen          int    `xml:"totallen"`           // 文件总长度
	AESKey            string `xml:"aeskey"`             // 加密密钥
	FileExt           string `xml:"fileext"`            // 文件扩展名
	OverwriteNewMsgID int64  `xml:"overwrite_newmsgid"` // 覆盖消息ID（需用int64处理大数）
}

func (d *DownloadHandler) DownloadFile(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用DownloadHandler->DownloadFile")

	var (
		req v1.DownloadFileWxRequest
		msg FileMsg
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("DownloadHandler->DownloadFile参数绑定失败")
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

	resp, err := d.sdk.DownloadFile(
		ctx, download.DownloadFileRequest{
			Appid:     req.AppId,
			AesKey:    msg.AppMsg.Attach.AESKey,
			FileType:  5,
			FileUrl:   msg.AppMsg.Attach.CDNAttachURL,
			TotalSize: msg.AppMsg.Attach.TotalLen,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用DownloadFile方法发送消息失败")
		response.Fail(ctx, errno.DownloadFileError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Err(err).Msg("调用DownloadFile方法发送消息失败")
		response.Fail(ctx, errno.DownloadFileError)
		return
	}

	filename := msg.AppMsg.Title
	path := filepath.Dir(execPath) + "/public/download/" + time.Now().Local().Format("20060102") + "/" + req.AppId + "/" + filename

	log.C(ctx).Info().Str("path", path).Msg("文件保存的目录")

	if err = help.DownloadBase64File(resp.Data.FileData, path); err != nil {
		log.C(ctx).Error().Err(err).Msg("下载文件失败")
		response.Fail(ctx, errno.DownloadImgError)
		return
	}

	response.Success(
		ctx, v1.DownloadFileWxResponse{
			FileUrl: path,
		},
	)
}

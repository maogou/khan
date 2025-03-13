package download

import (
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

func (d *DownloadHandler) DownloadCdn(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用DownloadHandler->DownloadCdn方法")

	var req v1.DownloadCdnRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("DownloadHandler->DownloadCdn方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	execPath, err := os.Executable()
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("获取程序路径失败")
		response.Fail(ctx, errno.ExecPathError)
		return
	}

	resp, err := d.sdk.DownloadCdn(
		ctx, download.DownloadCdnRequest{
			Appid:     req.AppId,
			AesKey:    req.AesKey,
			FileType:  req.Type,
			FileUrl:   req.FileId,
			Totalsize: req.TotalSize,
		},
	)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("调用DownloadCdn方法下载资源失败")
		response.Fail(ctx, errno.DownloadCdnError)
		return
	}

	if resp.Ret != 0 {
		log.C(ctx).Error().Any("resp", resp).Msg("resp.Ret != 0->调用DownloadCdn方法下载资源失败")
		response.Fail(ctx, errno.DownloadCdnError)
		return
	}

	filename := xid.New().String() + "." + req.Suffix
	path := filepath.Dir(execPath) + "/public/download/" + time.Now().Local().Format("20060102") + "/" + req.AppId + "/" + filename

	log.C(ctx).Info().Str("path", path).Msg("cdn资源保存的目录")

	if err = help.DownloadBase64File(resp.Data.FileData, path); err != nil {
		log.C(ctx).Error().Err(err).Msg("下载图片失败")
		response.Fail(ctx, errno.DownloadImgError)
		return
	}

	parts := strings.Split(path, "public/download")
	fileUrl := "http://localhost:" + strconv.Itoa(d.sdk.Config().Port) + parts[1]

	response.Success(
		ctx, v1.DownloadCdnResponse{
			FileUrl: fileUrl,
		},
	)
}

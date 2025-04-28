package download

import (
	v1 "maogou/khan/api/khan/v1"
	"maogou/khan/internal/pkg/errno"
	"maogou/khan/internal/pkg/log"
	"maogou/khan/internal/pkg/response"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (d *DownloadHandler) StaticFile(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用DownloadHandler->StaticFile方法")

	var req v1.DownloadFileRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		log.C(ctx).Error().Err(err).Msg("DownloadHandler->StaticFile方法参数绑定失败")
		response.Fail(ctx, errno.ValidateError)
		return
	}

	fileName := filepath.Base(req.File)
	filePath := filepath.Join("public", "download", req.File)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.C(ctx).Error().Err(err).Msg("文件不存在")
		response.Fail(ctx, errno.DownloadFileNotExistError)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.C(ctx).Error().Msg("文件打开失败")
		response.Fail(ctx, errno.DownloadFileOpenError)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.C(ctx).Error().Msg("获取文件信息失败")
		response.Fail(ctx, errno.DownloadFileStatError)
		return
	}

	ctx.Header("Content-Disposition", "attachment; filename="+fileName)

	http.ServeContent(ctx.Writer, ctx.Request, fileName, fileInfo.ModTime(), file)
}

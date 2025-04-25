package setting

import (
	"io"
	"smallBot/internal/constant"
	"smallBot/internal/pkg/errno"
	"smallBot/internal/pkg/help"
	"smallBot/internal/pkg/log"
	"smallBot/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (s *SettingHandler) SetLicense(ctx *gin.Context) {
	log.C(ctx).Info().Msg("调用SettingHandler->SetLicense方法")

	file, err := ctx.FormFile("file")

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("文件上传失败")
		response.Fail(ctx, errno.UploadFileError)
		return
	}

	fileName := file.Filename
	//读取文件内容
	src, err := file.Open()
	if err != nil {
		log.C(ctx).Error().Err(err).Msg("文件打开失败")
		response.Fail(ctx, errno.UploadFileError)
		return
	}

	defer src.Close()

	//把src中的内容读取到content变量中
	lByte, err := io.ReadAll(src)

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("文件读取失败")
		response.Fail(ctx, errno.UploadFileError)
		return
	}

	if len(lByte) == 0 {
		log.C(ctx).Error().Msg("文件内容为空")
		response.Fail(ctx, errno.UploadFileEmptyError)
		return
	}

	if err = s.sdk.Rdb().Set(ctx, constant.LicenseKey, fileName, 0).Err(); err != nil {
		log.C(ctx).Error().Err(err).Msg("设置license-key缓存失败")
		response.Fail(ctx, errno.SetLicenseKeyCacheError)
		return
	}

	if err = s.sdk.Rdb().Set(ctx, constant.License, string(lByte), 0).Err(); err != nil {
		log.C(ctx).Error().Err(err).Msg("设置license缓存失败")
		response.Fail(ctx, errno.SetLicenseCacheError)
		return
	}

	p, err := help.Permission(ctx, s.sdk.Rdb())

	if err != nil {
		log.C(ctx).Error().Err(err).Msg("解析license失败")
		response.Fail(ctx, errno.GetLicenseError)
		return
	}

	response.Success(ctx, p)
}

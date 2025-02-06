package help

import (
	"encoding/base64"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

func DownloadBase64File(base64Str, path string) error {
	dir := filepath.Dir(path)
	if err := Mkdir(dir); err != nil {
		return err
	}

	content, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return err
	}

	log.Info().Str("content", string(content)).Msg("图片下载完成")

	if err = os.WriteFile(path, content, 0644); err != nil {
		return err
	}

	return nil
}

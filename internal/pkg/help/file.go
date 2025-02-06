package help

import (
	"encoding/base64"
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

	if err = os.WriteFile(path, content, 0644); err != nil {
		return err
	}

	return nil
}

package help

import (
	"encoding/base64"
	"os"
	"path/filepath"

	"github.com/go-resty/resty/v2"
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

func GetFileBase64(client *resty.Client, fileUrl string) (string, error) {
	resp, err := client.R().Get(fileUrl)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(resp.Body()), nil
}

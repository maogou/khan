package download

import "smallBot/internal/sdk/khan"

type DownloadHandler struct {
	sdk *khan.Khan
}

func NewDownloadHandler(sdk *khan.Khan) *DownloadHandler {
	return &DownloadHandler{
		sdk: sdk,
	}
}

package v1

type DownloadFileRequest struct {
	File string `form:"file" binding:"required"`
}

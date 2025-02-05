package v1

type SnsUploadImgRequest struct {
	AppId   string   `json:"appId" binding:"required"`
	ImgUrls []string `json:"imgUrls" binding:"required,max=9"`
}

type SnsUploadImgItem struct {
	FileUrl  string `json:"fileUrl"`
	ThumbUrl string `json:"thumbUrl"`
	FileMd5  string `json:"fileMd5"`
	Length   int    `json:"length"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
}

type SnsUploadImgResponse struct {
	Data []SnsUploadImgItem `json:"data"`
}

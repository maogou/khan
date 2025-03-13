package v1

type DownloadFileRequest struct {
	File string `form:"file" binding:"required"`
}

type DownloadEmojiRequest struct {
	AppId    string `json:"appId" binding:"required"`
	EmojiMd5 string `json:"emojiMd5" binding:"required"`
}

type DownloadEmojiResponse struct {
	Url string `json:"url"`
}

type DownloadCdnRequest struct {
	AppId     string `json:"appId" binding:"required"`
	AesKey    string `json:"aesKey" binding:"required"`
	TotalSize int    `json:"totalSize" binding:"required"`
	Type      int    `json:"type" binding:"required,oneof=1 2 3 4 5"`
	FileId    string `json:"fileId" binding:"required"`
	Suffix    string `json:"suffix" binding:"required"`
}

type DownloadCdnResponse struct {
	FileUrl string `json:"fileUrl"`
}

type DownloadVideoRequest struct {
	AppId string `json:"appId" binding:"required"`
	Xml   string `json:"xml" binding:"required"`
}

type DownloadVideoResponse struct {
	FileUrl string `json:"fileUrl"`
}

type DownloadFileWxRequest struct {
	AppId string `json:"appId" binding:"required"`
	Xml   string `json:"xml" binding:"required"`
}

type DownloadFileWxResponse struct {
	FileUrl string `json:"fileUrl"`
}

type DownloadVoiceRequest struct {
	AppId string `json:"appId" binding:"required"`
	MsgId int    `json:"msgId" binding:"required"`
	Xml   string `json:"xml" binding:"required"`
}

type DownloadVoiceResponse struct {
	FileUrl string `json:"fileUrl"`
}

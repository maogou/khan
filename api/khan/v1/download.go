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

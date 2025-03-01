package v1

type PostTextRequest struct {
	AppId   string `json:"appId" binding:"required"`
	ToWxid  string `json:"toWxid" binding:"required"`
	Ats     string `json:"ats"`
	Content string `json:"content" binding:"required"`
}

type PostTextResponse struct {
	ToWxid     string `json:"toWxid"`
	CreateTime int    `json:"createTime"`
	MsgId      int    `json:"msgId"`
	NewMsgId   int64  `json:"newMsgId"`
	Type       int    `json:"type"`
}

type PostImageRequest struct {
	AppId  string `json:"appId" binding:"required"`
	ToWxid string `json:"toWxid" binding:"required"`
	ImgUrl string `json:"imgUrl" binding:"required"`
}

type PostImageResponse struct {
	ToWxid     string      `json:"toWxid"`
	CreateTime int         `json:"createTime"`
	MsgId      int         `json:"msgId"`
	NewMsgId   int64       `json:"newMsgId"`
	Type       interface{} `json:"type"`
	AesKey     string      `json:"aesKey"`
	FileId     string      `json:"fileId"`
	Length     int         `json:"length"`
	Width      int         `json:"width"`
	Height     int         `json:"height"`
	Md5        string      `json:"md5"`
}

type PostFileRequest struct {
	AppId    string `json:"appId"  binding:"required"`
	ToWxid   string `json:"toWxid" binding:"required"`
	FileName string `json:"fileName" binding:"required"`
	FileUrl  string `json:"fileUrl" binding:"required"`
}

type PostFileResponse struct {
	ToWxid     string `json:"toWxid"`
	CreateTime int    `json:"createTime"`
	MsgId      int    `json:"msgId"`
	NewMsgId   int64  `json:"newMsgId"`
	Type       int    `json:"type"`
}

type PostLinkRequest struct {
	AppId    string `json:"appId" binding:"required"`
	ToWxid   string `json:"toWxid" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Desc     string `json:"desc" binding:"required"`
	LinkUrl  string `json:"linkUrl" binding:"required"`
	ThumbUrl string `json:"thumbUrl" binding:"required"`
}

type PostLinkResponse struct {
	ToWxid     string `json:"toWxid"`
	CreateTime int    `json:"createTime"`
	MsgId      int    `json:"msgId"`
	NewMsgId   int64  `json:"newMsgId"`
	Type       int    `json:"type"`
}

type PostVoiceRequest struct {
	AppId         string `json:"appId" binding:"required"`
	ToWxid        string `json:"toWxid" binding:"required"`
	VoiceUrl      string `json:"voiceUrl" binding:"required"`
	VoiceDuration int    `json:"voiceDuration" binding:"required"`
}

type PostVoiceResponse struct {
	ToWxid     string `json:"toWxid"`
	CreateTime int    `json:"createTime"`
	MsgId      int    `json:"msgId"`
	NewMsgId   int64  `json:"newMsgId"`
	Type       int    `json:"type"`
}

type PostVideoRequest struct {
	AppId         string `json:"appId" binding:"required"`
	ToWxid        string `json:"toWxid" binding:"required"`
	VideoUrl      string `json:"videoUrl" binding:"required"`
	ThumbUrl      string `json:"thumbUrl" binding:"required"`
	VideoDuration int    `json:"videoDuration" binding:"required"`
}

type PostVideoResponse struct {
	ToWxid     string      `json:"toWxid"`
	CreateTime interface{} `json:"createTime"`
	MsgId      int         `json:"msgId"`
	NewMsgId   int64       `json:"newMsgId"`
	Type       int         `json:"type"`
	AesKey     string      `json:"aesKey"`
	FileId     string      `json:"fileId"`
	Length     int         `json:"length"`
}

type PostNameCardRequest struct {
	AppId        string `json:"appId" binding:"required"`
	ToWxid       string `json:"toWxid" binding:"required"`
	NickName     string `json:"nickName" binding:"required"`
	NameCardWxid string `json:"nameCardWxid" binding:"required"`
}

type PostNameCardResponse struct {
	ToWxid     string `json:"toWxid"`
	CreateTime int    `json:"createTime"`
	MsgId      int    `json:"msgId"`
	NewMsgId   int64  `json:"newMsgId"`
	Type       int    `json:"type"`
}

type PostEmojiRequest struct {
	AppId     string `json:"appId" binding:"required"`
	ToWxid    string `json:"toWxid" binding:"required"`
	EmojiMd5  string `json:"emojiMd5" binding:"required"`
	EmojiSize int    `json:"emojiSize" binding:"required"`
}

type PostEmojiResponse struct {
	ToWxid     string `json:"toWxid"`
	CreateTime int64  `json:"createTime"`
	MsgId      int    `json:"msgId"`
	NewMsgId   int64  `json:"newMsgId"`
	Type       int    `json:"type"`
}

type DownloadImgRequest struct {
	AppId string `json:"appId" binding:"required"`
	Type  int    `json:"type" binding:"required,oneof=1 2 3"`
	Xml   string `json:"xml" binding:"required"`
}

type DownloadImgResponse struct {
	FileUrl string `json:"fileUrl"`
}

type RevokeMsgRequest struct {
	AppId      string `json:"appId" binding:"required"`
	ToWxid     string `json:"toWxid" binding:"required"`
	NewMsgId   int64  `json:"newMsgId" binding:"required"`
	CreateTime int64  `json:"createTime" binding:"required"`
}

type ForwardMiniAppRequest struct {
	AppId       string `json:"appId" binding:"required"`
	ToWxid      string `json:"toWxid" binding:"required"`
	Xml         string `json:"xml" binding:"required"`
	CoverImgUrl string `json:"coverImgUrl" binding:"required"`
}

type ForwardMiniAppResponse struct {
	ToWxid     string `json:"toWxid"`
	CreateTime int    `json:"createTime"`
	MsgId      int    `json:"msgId"`
	NewMsgId   int64  `json:"newMsgId"`
	Type       int    `json:"type"`
}

type ForwardUrlRequest struct {
	AppId  string `json:"appId" binding:"required"`
	ToWxid string `json:"toWxid" binding:"required"`
	Xml    string `json:"xml" binding:"required"`
}

type ForwardUrlResponse struct {
	ToWxid     string `json:"toWxid"`
	CreateTime int    `json:"createTime"`
	MsgId      int    `json:"msgId"`
	NewMsgId   int64  `json:"newMsgId"`
	Type       int    `json:"type"`
}

type ForwardVideoRequest struct {
	AppId  string `json:"appId" binding:"required"`
	ToWxid string `json:"toWxid" binding:"required"`
	Xml    string `json:"xml" binding:"required"`
}

type ForwardVideoResponse struct {
	ToWxid     string `json:"toWxid"`
	CreateTime int64  `json:"createTime"`
	MsgId      int    `json:"msgId"`
	NewMsgId   int64  `json:"newMsgId"`
	AesKey     string `json:"aesKey"`
}

type ForwardImageRequest struct {
	AppId  string `json:"appId" binding:"required"`
	ToWxid string `json:"toWxid" binding:"required"`
	Xml    string `json:"xml" binding:"required"`
}

type ForwardImageResponse struct {
	ToWxid     string `json:"toWxid"`
	CreateTime int    `json:"createTime"`
	MsgId      int    `json:"msgId"`
	NewMsgId   int64  `json:"newMsgId"`
	AesKey     string `json:"aesKey"`
	FileId     string `json:"fileId"`
}

type ForwardFileRequest struct {
	AppId  string `json:"appId" binding:"required"`
	ToWxid string `json:"toWxid" binding:"required"`
	Xml    string `json:"xml" binding:"required"`
}

type ForwardFileResponse struct {
	ToWxid     string `json:"toWxid"`
	CreateTime int    `json:"createTime"`
	MsgId      int    `json:"msgId"`
	NewMsgId   int64  `json:"newMsgId"`
	Type       int    `json:"type"`
}

type SendMiniAppRequest struct {
	AppId       string `json:"appId" binding:"required"`
	ToWxid      string `json:"toWxid" binding:"required"`
	MiniAppId   string `json:"miniAppId" binding:"required"`
	UserName    string `json:"userName" binding:"required"`
	Title       string `json:"title" binding:"required"`
	CoverImgUrl string `json:"coverImgUrl" binding:"required"`
	PagePath    string `json:"pagePath" binding:"required"`
	DisplayName string `json:"displayName" binding:"required"`
}

type SendMiniAppResponse struct {
	ToWxid     string `json:"toWxid"`
	CreateTime int    `json:"createTime"`
	MsgId      int    `json:"msgId"`
	NewMsgId   int64  `json:"newMsgId"`
	Type       int    `json:"type"`
}

type PostAppMsgRequest struct {
	AppId  string `json:"appId" binding:"required"`
	ToWxid string `json:"toWxid" binding:"required"`
	Appmsg string `json:"appmsg" binding:"required"`
}

type PostAppMsgResponse struct {
	ToWxid     string `json:"toWxid"`
	CreateTime int    `json:"createTime"`
	MsgId      int    `json:"msgId"`
	NewMsgId   int64  `json:"newMsgId"`
	Type       int    `json:"type"`
}

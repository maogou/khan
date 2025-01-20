package v1

type PostTextRequest struct {
	AppId   string `json:"appId"`
	ToWxid  string `json:"toWxid"`
	Ats     string `json:"ats"`
	Content string `json:"content"`
}

type PostTextResponse struct {
	ToWxid     string `json:"toWxid"`
	CreateTime int    `json:"createTime"`
	MsgId      int    `json:"msgId"`
	NewMsgId   int64  `json:"newMsgId"`
	Type       int    `json:"type"`
}

type PostImageRequest struct {
	AppId  string `json:"appId" validate:"required"`
	ToWxid string `json:"toWxid" validate:"required"`
	ImgUrl string `json:"imgUrl" validate:"required"`
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
	AppId    string `json:"appId"  validate:"required"`
	ToWxid   string `json:"toWxid" validate:"required"`
	FileName string `json:"fileName" validate:"required"`
	FileUrl  string `json:"fileUrl" validate:"required"`
}

type PostFileResponse struct {
	ToWxid     string `json:"toWxid"`
	CreateTime int    `json:"createTime"`
	MsgId      int    `json:"msgId"`
	NewMsgId   int64  `json:"newMsgId"`
	Type       int    `json:"type"`
}

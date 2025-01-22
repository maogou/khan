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

type PersonalProfileRequest struct {
	AppId   string `json:"appId" binding:"required"`
	ProxyIp string `json:"proxyIp"`
}

type PersonalProfileResponse struct {
	Alias           string `json:"alias"`
	Wxid            string `json:"wxid"`
	NickName        string `json:"nickName"`
	Mobile          string `json:"mobile"`
	Uin             int    `json:"uin"`
	Sex             int    `json:"sex"`
	Province        string `json:"province"`
	City            string `json:"city"`
	Signature       string `json:"signature"`
	Country         string `json:"country"`
	BigHeadImgUrl   string `json:"bigHeadImgUrl"`
	SmallHeadImgUrl string `json:"smallHeadImgUrl"`
	RegCountry      string `json:"regCountry"`
	SnsBgImg        string `json:"snsBgImg"`
}

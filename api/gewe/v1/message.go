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

type PersonalQrcodeRequest struct {
	AppId   string `json:"appId" binding:"required"`
	ProxyIp string `json:"proxyIp"`
}

type PersonalQrcodeResponse struct {
	QrCode string `json:"qrCode"`
}

type PersonalSafetyRequest struct {
	AppId   string `json:"appId" binding:"required"`
	ProxyIp string `json:"proxyIp"`
}

type PersonalSafetyItem struct {
	Uuid       string `json:"uuid"`
	DeviceName string `json:"deviceName"`
	DeviceType string `json:"deviceType"`
	LastTime   int    `json:"lastTime"`
}

type PersonalSafetResponse struct {
	List []PersonalSafetyItem `json:"list"`
}

type PersonalPrivacySettingRequest struct {
	AppId  string `json:"appId" binding:"required"`
	Open   bool   `json:"open" binding:"required"`
	Option int    `json:"option" binding:"required"`
}

type PersonalUploadHeadImgRequest struct {
	AppId      string `json:"appId" binding:"required"`
	HeadImgUrl string `json:"headImgUrl" binding:"required"`
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

type ContactListRequest struct {
	AppId string `json:"appId" binding:"required"`
}

type ContactListResponse struct {
	Friends   []string `json:"friends"`
	Chatrooms []string `json:"chatrooms"`
	Ghs       []string `json:"ghs"`
}

type ContactSearchRequest struct {
	AppId        string `json:"appId" binding:"required"`
	ContactsInfo string `json:"contactsInfo" binding:"required"`
}

type ContactSearchResponse struct {
	V3              string `json:"v3"`
	NickName        string `json:"nickName"`
	Sex             int    `json:"sex"`
	Signature       string `json:"signature"`
	BigHeadImgUrl   string `json:"bigHeadImgUrl"`
	SmallHeadImgUrl string `json:"smallHeadImgUrl"`
	V4              string `json:"v4"`
}

type ContactAddRequest struct {
	AppId   string `json:"appId" binding:"required"`
	Scene   int    `json:"scene" binding:"required"`
	Content string `json:"content" binding:"required"`
	V4      string `json:"v4" binding:"required"`
	V3      string `json:"v3" binding:"required"`
	Option  int    `json:"option" binding:"required"`
}

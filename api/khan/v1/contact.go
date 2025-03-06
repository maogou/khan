package v1

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

type ContactDelRequest struct {
	AppId string `json:"appId" binding:"required"`
	Wxid  string `json:"wxid" binding:"required"`
}

type ContactInfoRequest struct {
	AppId string   `json:"appId" binding:"required"`
	Wxids []string `json:"wxids" binding:"required,max=100"`
}

type ContactInfoItem struct {
	UserName        string `json:"userName"`
	NickName        string `json:"nickName"`
	PyInitial       string `json:"pyInitial"`
	QuanPin         string `json:"quanPin"`
	Sex             int    `json:"sex"`
	Remark          string `json:"remark"`
	RemarkPyInitial string `json:"remarkPyInitial"`
	RemarkQuanPin   string `json:"remarkQuanPin"`
	Signature       string `json:"signature"`
	Alias           string `json:"alias"`
	SnsBgImg        string `json:"snsBgImg"`
	Country         string `json:"country"`
	BigHeadImgUrl   string `json:"bigHeadImgUrl"`
	SmallHeadImgUrl string `json:"smallHeadImgUrl"`
	Description     string `json:"description"`
	CardImgUrl      string `json:"cardImgUrl"`
	LabelList       string `json:"labelList"`
	Province        string `json:"province"`
	City            string `json:"city"`
	PhoneNumList    string `json:"phoneNumList"`
}

type ContactDetailRequest struct {
	AppId string   `json:"appId" binding:"required"`
	Wxids []string `json:"wxids" binding:"required,max=20"`
}

type ContactDetailItem struct {
	UserName        string      `json:"userName"`
	NickName        string      `json:"nickName"`
	PyInitial       interface{} `json:"pyInitial"`
	QuanPin         string      `json:"quanPin"`
	Sex             int         `json:"sex"`
	Remark          interface{} `json:"remark"`
	RemarkPyInitial interface{} `json:"remarkPyInitial"`
	RemarkQuanPin   interface{} `json:"remarkQuanPin"`
	Signature       string      `json:"signature"`
	Alias           string      `json:"alias"`
	SnsBgImg        string      `json:"snsBgImg"`
	Country         string      `json:"country"`
	BigHeadImgUrl   string      `json:"bigHeadImgUrl"`
	SmallHeadImgUrl string      `json:"smallHeadImgUrl"`
	Description     interface{} `json:"description"`
	CardImgUrl      interface{} `json:"cardImgUrl"`
	LabelList       interface{} `json:"labelList"`
	Province        interface{} `json:"province"`
	City            interface{} `json:"city"`
	PhoneNumList    interface{} `json:"phoneNumList"`
}

type ContactSetOnlyChatRequest struct {
	AppId    string `json:"appId" binding:"required"`
	Wxid     string `json:"wxid" binding:"required"`
	OnlyChat bool   `json:"onlyChat"`
}

type ContactSetRemarkResquest struct {
	AppId  string `json:"appId" binding:"required"`
	Wxid   string `json:"wxid" binding:"required"`
	Remark string `json:"remark" binding:"required"`
}

type ContactUploadMobileRequest struct {
	AppId  string   `json:"appId" binding:"required"`
	Phones []string `json:"phones" binding:"required"`
	OpType int      `json:"opType" binding:"required,oneof=1 2"`
}

type ContactGetMobileRequest struct {
	AppId  string   `json:"appId" binding:"required"`
	Phones []string `json:"phones"`
}
type ContactGetMobileItem struct {
	UserName        string `json:"userName"`
	NickName        string `json:"nickName"`
	Sex             int    `json:"sex"`
	PhoneMd5        string `json:"phoneMd5"`
	Signature       string `json:"signature"`
	Country         string `json:"country"`
	BigHeadImgUrl   string `json:"bigHeadImgUrl"`
	SmallHeadImgUrl string `json:"smallHeadImgUrl"`
	Province        string `json:"province"`
	City            string `json:"city"`
	PersonalCard    int    `json:"personalCard"`
}

package v1

import "encoding/xml"

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

type SnsDetailRequest struct {
	AppId string `json:"appId" binding:"required"`
	SnsId string `json:"snsId" binding:"required"`
}

type SnsLikeListItem struct {
	UserName   string `json:"userName"`
	NickName   string `json:"nickName"`
	Source     int    `json:"source"`
	Type       int    `json:"type"`
	CreateTime int64  `json:"createTime"`
}

type SnsCommentListItem struct {
	UserName        string `json:"userName"`
	NickName        string `json:"nickName"`
	Source          int    `json:"source"`
	Type            int    `json:"type"`
	Content         string `json:"content"`
	CreateTime      int64  `json:"createTime"`
	CommentId       int    `json:"commentId"`
	ReplyCommentId  int    `json:"replyCommentId"`
	ReplyCommentId2 int    `json:"replyCommentId2"`
	IsNotRichText   int    `json:"isNotRichText"`
}

type SnsWithUserListItem struct {
	Username        string `json:"Username"`
	Source          int    `json:"Source"`
	Type            int    `json:"Type"`
	CreateTime      int64  `json:"CreateTime"`
	CommentId       int    `json:"CommentId"`
	ReplyCommentId  int    `json:"ReplyCommentId"`
	IsNotRichText   int    `json:"IsNotRichText"`
	ReplyCommentId2 int    `json:"ReplyCommentId2"`
	CommentId2      int    `json:"CommentId2"`
	DeleteFlag      int    `json:"DeleteFlag"`
	CommentFlag     int    `json:"CommentFlag"`
}

type SnsDetailResponse struct {
	Id            int64                 `json:"id"`
	UserName      string                `json:"userName"`
	NickName      string                `json:"nickName"`
	CreateTime    int64                 `json:"createTime"`
	SnsXml        string                `json:"snsXml"`
	LikeCount     int                   `json:"likeCount"`
	LikeList      []SnsLikeListItem     `json:"likeList"`
	CommentCount  int                   `json:"commentCount"`
	CommentList   []SnsCommentListItem  `json:"commentList"`
	WithUserCount int                   `json:"withUserCount"`
	WithUserList  []SnsWithUserListItem `json:"withUserList"`
}

type SnsForwardRequest struct {
	AppId        string   `json:"appId" binding:"required"`
	AllowWxIds   []string `json:"allowWxIds"`
	AtWxIds      []string `json:"atWxIds"`
	DisableWxIds []string `json:"disableWxIds"`
	SnsXml       string   `json:"snsXml" binding:"required"`
	Privacy      bool     `json:"privacy"`
}

type SnsForwardResponse struct {
	Id         float64 `json:"id"`
	UserName   string  `json:"userName"`
	NickName   string  `json:"nickName"`
	CreateTime int64   `json:"createTime"`
}

type SnsPageItem struct {
	Id            string                `json:"id"`
	UserName      string                `json:"userName"`
	NickName      string                `json:"nickName"`
	CreateTime    int64                 `json:"createTime"`
	SnsXml        string                `json:"snsXml"`
	LikeCount     int                   `json:"likeCount"`
	LikeList      []SnsLikeListItem     `json:"likeList"`
	CommentCount  int                   `json:"commentCount"`
	CommentList   []SnsCommentListItem  `json:"commentList"`
	WithUserCount int                   `json:"withUserCount"`
	WithUserList  []SnsWithUserListItem `json:"withUserList"`
}

type SnsFriendPageRequest struct {
	AppId        string `json:"appId" binding:"required"`
	MaxId        int    `json:"maxId"`
	Decrypt      bool   `json:"decrypt"`
	Wxid         string `json:"wxid" binding:"required"`
	FirstPageMd5 string `json:"firstPageMd5"`
}

type SnsFriendPageResponse struct {
	FirstPageMd5 string        `json:"firstPageMd5"`
	MaxId        string        `json:"maxId"`
	SnsCount     int           `json:"snsCount"`
	RequestTime  int           `json:"requestTime"`
	SnsList      []SnsPageItem `json:"snsList"`
}

type SnsMyselfPageRequest struct {
	AppId        string `json:"appId" binding:"required"`
	MaxId        int    `json:"maxId"`
	Decrypt      bool   `json:"decrypt"`
	FirstPageMd5 string `json:"firstPageMd5"`
}

type SnsMyselfPageResponse struct {
	FirstPageMd5 string        `json:"firstPageMd5"`
	MaxId        string        `json:"maxId"`
	SnsCount     int           `json:"snsCount"`
	RequestTime  int           `json:"requestTime"`
	SnsList      []SnsPageItem `json:"snsList"`
}

type SnsLikeRequest struct {
	AppId string `json:"appId" binding:"required"`
	SnsId string `json:"snsId" binding:"required"`
	Wxid  string `json:"wxid" binding:"required"`
}

type SnsCancelLikeRequest struct {
	AppId string `json:"appId" binding:"required"`
	SnsId string `json:"snsId" binding:"required"`
}

type SnsCommentRequest struct {
	AppId     string `json:"appId" binding:"required"`
	SnsId     string `json:"snsId" binding:"required"`
	Wxid      string `json:"wxid" binding:"required"`
	Content   string `json:"content" binding:"required"`
	CommentId int    `json:"commentId"`
}

type SnsDelCommentRequest struct {
	AppId     string `json:"appId" binding:"required"`
	SnsId     string `json:"snsId" binding:"required"`
	CommentId int    `json:"commentId" binding:"required"`
}

type SnsSetPrivacyScopeRequest struct {
	AppId  string `json:"appId" binding:"required"`
	Option int    `json:"option" binding:"required,oneof=1 2 3 4"`
}

type SnsStrangerVisibilityEnabledRequest struct {
	AppId   string `json:"appId" binding:"required"`
	Enabled bool   `json:"enabled"`
}

type SnsSetPrivacyRequest struct {
	AppId string `json:"appId" binding:"required"`
	SnsId string `json:"snsId" binding:"required"`
	Open  bool   `json:"open"`
}

type SnsDeleteRequest struct {
	AppId string `json:"appId"  binding:"required"`
	SnsId string `json:"snsId" binding:"required"`
}

type SnsUploadVideoRequest struct {
	AppId    string `json:"appId" binding:"required"`
	ThumbUrl string `json:"thumbUrl" binding:"required"`
	VideoUrl string `json:"videoUrl" binding:"required"`
}

type SnsUploadVideoResponse struct {
	FileUrl  string `json:"fileUrl"`
	ThumbUrl string `json:"thumbUrl"`
	FileMd5  string `json:"fileMd5"`
	Length   int    `json:"length"`
}

type SnsSendTextRequest struct {
	AppId        string   `json:"appId"`
	AllowWxIds   []string `json:"allowWxIds"`
	AtWxIds      []string `json:"atWxIds"`
	DisableWxIds []string `json:"disableWxIds"`
	Content      string   `json:"content"`
	Privacy      bool     `json:"privacy"`
}

type SnsSendTextResponse struct {
	Id         float64 `json:"id"`
	UserName   string  `json:"userName"`
	NickName   string  `json:"nickName"`
	CreateTime int64   `json:"createTime"`
}

type SnsSendImageItem struct {
	FileUrl  string `json:"fileUrl" binding:"required"`
	ThumbUrl string `json:"thumbUrl" binding:"required"`
	FileMd5  string `json:"fileMd5" binding:"required"`
	Length   int    `json:"length"`
	Width    int    `json:"width" binding:"required"`
	Height   int    `json:"height" binding:"required"`
}

type SnsSendImageRequest struct {
	AppId        string             `json:"appId" binding:"required"`
	AllowWxIds   []string           `json:"allowWxIds"`
	AtWxIds      []string           `json:"atWxIds"`
	DisableWxIds []string           `json:"disableWxIds"`
	Content      string             `json:"content"`
	ImgInfos     []SnsSendImageItem `json:"imgInfos" binding:"required"`
	Privacy      bool               `json:"privacy"`
}

type SnsSendImageResponse struct {
	Id         float64 `json:"id"`
	UserName   string  `json:"userName"`
	NickName   string  `json:"nickName"`
	CreateTime int64   `json:"createTime"`
}

type SnsSendUrlRequest struct {
	AppId        string   `json:"appId" binding:"required"`
	AllowWxIds   []string `json:"allowWxIds"`
	AtWxIds      []string `json:"atWxIds"`
	DisableWxIds []string `json:"disableWxIds"`
	Content      string   `json:"content"`
	Description  string   `json:"description" binding:"required"`
	Title        string   `json:"title" binding:"required"`
	LinkUrl      string   `json:"linkUrl" binding:"required"`
	ThumbUrl     string   `json:"thumbUrl" binding:"required"`
	Privacy      bool     `json:"privacy"`
}

type SnsSendUrlResponse struct {
	Id         float64 `json:"id"`
	UserName   string  `json:"userName"`
	NickName   string  `json:"nickName"`
	CreateTime int64   `json:"createTime"`
}

type TimelineObject struct {
	XMLName             xml.Name      `xml:"TimelineObject"`
	Id                  string        `xml:"id"`
	Username            string        `xml:"username"`
	CreateTime          string        `xml:"createTime"`
	ContentDesc         string        `xml:"contentDesc"`
	ContentDescShowType string        `xml:"contentDescShowType"`
	ContentDescScene    string        `xml:"contentDescScene"`
	Private             string        `xml:"private"`
	SightFolded         string        `xml:"sightFolded"`
	ShowFlag            string        `xml:"showFlag"`
	AppInfo             AppInfo       `xml:"appInfo"`
	SourceUserName      string        `xml:"sourceUserName"`
	SourceNickName      string        `xml:"sourceNickName"`
	StatisticsData      string        `xml:"statisticsData"`
	StatExtStr          string        `xml:"statExtStr"`
	ContentObject       ContentObject `xml:"ContentObject"`
	ActionInfo          ActionInfo    `xml:"actionInfo"`
	Location            Location      `xml:"location"`
	PublicUserName      string        `xml:"publicUserName"`
	Streamvideo         Streamvideo   `xml:"streamvideo"`
}

type AppInfo struct {
	Id            string `xml:"id"`
	Version       string `xml:"version"`
	AppName       string `xml:"appName"`
	InstallUrl    string `xml:"installUrl"`
	FromUrl       string `xml:"fromUrl"`
	IsForceUpdate string `xml:"isForceUpdate"`
	IsHidden      string `xml:"isHidden"`
}

type ContentObject struct {
	ContentStyle int       `xml:"contentStyle"`
	Title        string    `xml:"title"`
	Description  string    `xml:"description"`
	MediaList    MediaList `xml:"mediaList"`
	ContentUrl   string    `xml:"contentUrl"`
}

type MediaList struct {
	Media []Media `xml:"media"`
}

type Media struct {
	Id          string    `xml:"id"`
	Type        string    `xml:"type"`
	Title       string    `xml:"title"`
	Description string    `xml:"description"`
	Private     string    `xml:"private"`
	UserData    string    `xml:"userData"`
	SubType     string    `xml:"subType"`
	VideoSize   VideoSize `xml:"videoSize"`
	Url         Url       `xml:"url"`
	Thumb       Thumb     `xml:"thumb"`
	Size        Size      `xml:"size"`
}

type VideoSize struct {
	Width  string `xml:"width,attr"`
	Height string `xml:"height,attr"`
}

type Url struct {
	Type     string `xml:"type,attr"`
	Md5      string `xml:"md5,attr"`
	Videomd5 string `xml:"videomd5,attr"`
	Text     string `xml:",chardata"`
}

type Thumb struct {
	Type string `xml:"type,attr"`
	Text string `xml:",chardata"`
}

type Size struct {
	Width     string `xml:"width,attr"`
	Height    string `xml:"height,attr"`
	TotalSize string `xml:"totalSize,attr"`
}

type ActionInfo struct {
	AppMsg AppMsg `xml:"appMsg"`
}

type AppMsg struct {
	MessageAction string `xml:"messageAction"`
}

type Location struct {
	PoiClassifyId   string `xml:"poiClassifyId,attr"`
	PoiName         string `xml:"poiName,attr"`
	PoiAddress      string `xml:"poiAddress,attr"`
	PoiClassifyType string `xml:"poiClassifyType,attr"`
	City            string `xml:"city,attr"`
}

type Streamvideo struct {
	Streamvideourl      string `xml:"streamvideourl"`
	Streamvideothumburl string `xml:"streamvideothumburl"`
	Streamvideoweburl   string `xml:"streamvideoweburl"`
}

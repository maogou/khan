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
	Id         string `json:"id"`
	UserName   string `json:"userName"`
	NickName   string `json:"nickName"`
	CreateTime int64  `json:"createTime"`
}

type SnsFriendPageRequest struct {
	AppId        string `json:"appId" binding:"required"`
	MaxId        int    `json:"maxId"`
	Decrypt      bool   `json:"decrypt"`
	Wxid         string `json:"wxid" binding:"required"`
	FirstPageMd5 string `json:"firstPageMd5"`
}

type SnsMyselfPageRequest struct {
	AppId        string `json:"appId" binding:"required"`
	MaxId        int    `json:"maxId"`
	Decrypt      bool   `json:"decrypt"`
	FirstPageMd5 string `json:"firstPageMd5"`
}

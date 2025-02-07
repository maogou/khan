package v1

type ChatRoomCreateRequest struct {
	AppId string   `json:"appId" binding:"required"`
	Wxids []string `json:"wxids" binding:"required,min=2,max=500"`
}

type ChatRoomCreateResponse struct {
	HeadImgBase64 string `json:"headImgBase64"`
	ChatroomId    string `json:"chatroomId"`
}

type ChatroomModifyNameRequest struct {
	AppId        string `json:"appId" binding:"required"`
	ChatroomName string `json:"chatroomName" binding:"required"`
	ChatroomId   string `json:"chatroomId" binding:"required"`
}

type ChatroomSetRemarkRequest struct {
	AppId          string `json:"appId" binding:"required"`
	ChatroomRemark string `json:"chatroomRemark" binding:"required"`
	ChatroomId     string `json:"chatroomId" binding:"required"`
}

type ChatroomModifyMyselfNameRequest struct {
	AppId      string `json:"appId" binding:"required"`
	NickName   string `json:"nickName" binding:"required"`
	ChatroomId string `json:"chatroomId" binding:"required"`
}

type ChatroomSetAnnouncementRequest struct {
	AppId      string `json:"appId" binding:"required"`
	ChatroomId string `json:"chatroomId" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

type ChatroomGetAnnouncementRequest struct {
	AppId      string `json:"appId" binding:"required"`
	ChatroomId string `json:"chatroomId" binding:"required"`
}

type ChatroomGetAnnouncementResponse struct {
	Announcement       string `json:"announcement"`
	AnnouncementEditor string `json:"announcementEditor"`
	PublishTime        int    `json:"publishTime"`
}

type ChatroomDetailRequest struct {
	AppId      string `json:"appId" binding:"required"`
	ChatroomId string `json:"chatroomId" binding:"required"`
}

type ChatroomMemberItem struct {
	Wxid            string `json:"wxid"`
	NickName        string `json:"nickName"`
	InviterUserName string `json:"inviterUserName"`
	MemberFlag      int    `json:"memberFlag"`
	DisplayName     string `json:"displayName"`
	BigHeadImgUrl   string `json:"bigHeadImgUrl"`
	SmallHeadImgUrl string `json:"smallHeadImgUrl"`
}

type ChatroomDetailResponse struct {
	ChatroomId      string               `json:"chatroomId"`
	UserName        string               `json:"userName"`
	NickName        string               `json:"nickName"`
	PyInitial       string               `json:"pyInitial"`
	QuanPin         string               `json:"quanPin"`
	Sex             int                  `json:"sex"`
	Remark          string               `json:"remark"`
	RemarkPyInitial string               `json:"remarkPyInitial"`
	RemarkQuanPin   string               `json:"remarkQuanPin"`
	ChatRoomNotify  int                  `json:"chatRoomNotify"`
	ChatRoomOwner   string               `json:"chatRoomOwner"`
	SmallHeadImgUrl string               `json:"smallHeadImgUrl"`
	MemberList      []ChatroomMemberItem `json:"memberList"`
}

type ChatroomMemberRequest struct {
	AppId      string `json:"appId" binding:"required"`
	ChatroomId string `json:"chatroomId" binding:"required"`
}

type ChatroomMemberResponse struct {
	MemberList    []ChatroomMemberItem `json:"memberList"`
	ChatroomOwner string               `json:"chatroomOwner"`
	AdminWxid     []string             `json:"adminWxid"`
}

type ChatroomTopPingRequest struct {
	AppId      string `json:"appId" binding:"required"`
	ChatroomId string `json:"chatroomId" binding:"required"`
	Top        bool   `json:"top"`
}

type ChatroomSetSilenceRequest struct {
	AppId      string `json:"appId" binding:"required"`
	ChatroomId string `json:"chatroomId" binding:"required"`
	Silence    bool   `json:"silence"`
}

type ChatroomQrcodeRequest struct {
	AppId      string `json:"appId" binding:"required"`
	ChatroomId string `json:"chatroomId" binding:"required"`
}

type ChatroomQrcodeResponse struct {
	QrBase64 string `json:"qrBase64"`
	QrTips   string `json:"qrTips"`
}

package transform

type ChatroomMemberRequest struct {
	Appid   string `json:"appid"`
	GroupId string `json:"group_id"`
}

type ChatroomMemberResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		ChatroomUserName string `json:"ChatroomUserName"`
		ServerVersion    int    `json:"ServerVersion"`
		NewChatroomData  struct {
			MemberCount    int `json:"MemberCount"`
			ChatRoomMember []struct {
				UserName           string `json:"UserName"`
				NickName           string `json:"NickName"`
				BigHeadImgUrl      string `json:"BigHeadImgUrl"`
				SmallHeadImgUrl    string `json:"SmallHeadImgUrl"`
				ChatroomMemberFlag int    `json:"ChatroomMemberFlag"`
				InviterUserName    string `json:"InviterUserName,omitempty"`
			} `json:"ChatRoomMember"`
			InfoMask int `json:"InfoMask"`
		} `json:"NewChatroomData"`
		ChatRoomOwner         string `json:"chatRoomOwner"`
		AllMemberCount        int    `json:"allMemberCount"`
		AllMemberUserNameList []struct {
			String string `json:"string"`
		} `json:"allMemberUserNameList"`
		AdminCount int `json:"adminCount"`
	} `json:"data"`
}

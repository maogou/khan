package chatroom

type ChatroomInviteRequest struct {
	Appid   string   `json:"appid"`
	Cause   string   `json:"cause"`
	GroupId string   `json:"group_id"`
	ToWxid  []string `json:"to_wxid"`
}

type ChatroomInviteResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
				String string `json:"string"`
			} `json:"errMsg"`
		} `json:"baseResponse"`
		MemberCount int `json:"MemberCount"`
		MemberList  []struct {
			MemberName struct {
				String string `json:"string"`
			} `json:"MemberName"`
			MemberStatus int `json:"MemberStatus"`
			NickName     struct {
				String string `json:"string"`
			} `json:"NickName"`
			PYInitial struct {
				String string `json:"string"`
			} `json:"PYInitial"`
			QuanPin struct {
				String string `json:"string"`
			} `json:"QuanPin"`
			Sex    int `json:"Sex"`
			Remark struct {
			} `json:"Remark"`
			RemarkPyinitial struct {
			} `json:"RemarkPyinitial"`
			RemarkQuanPin struct {
			} `json:"RemarkQuanPin"`
			ContactType  int    `json:"ContactType"`
			Province     string `json:"Province"`
			Signature    string `json:"Signature"`
			PersonalCard int    `json:"PersonalCard"`
			VerifyFlag   int    `json:"VerifyFlag"`
			Country      string `json:"Country"`
		} `json:"MemberList"`
	} `json:"data"`
}

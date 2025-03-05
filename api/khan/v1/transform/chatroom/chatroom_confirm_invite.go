package chatroom

type ChatroomConfirInviteRequest struct {
	Appid       string   `json:"appid"`
	GroupId     string   `json:"group_id"`
	InviterWxid string   `json:"inviterWxid"`
	Msgid       int64    `json:"msgid"`
	Ticket      string   `json:"ticket"`
	ToWxidList  []string `json:"to_wxid_list"`
}

type ChatroomConfirInviteResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
				String string `json:"string"`
			} `json:"errMsg"`
		} `json:"baseResponse"`
	} `json:"data"`
}

package chatroom

type ChatroomDeleteMebmerRequest struct {
	Appid      string   `json:"appid"`
	GroupId    string   `json:"group_id"`
	ToWxidList []string `json:"to_wxid_list"`
}

type ChatroomDeleteMebmerResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"baseResponse"`
		MemberCount int `json:"MemberCount"`
		MemberList  []struct {
			MemberName struct {
				String string `json:"string"`
			} `json:"MemberName"`
		} `json:"MemberList"`
	} `json:"data"`
}

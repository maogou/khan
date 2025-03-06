package chatroom

type ChatroomAdminOperateRequest struct {
	Appid      string   `json:"appid"`
	GroupId    string   `json:"group_id"`
	ToWxidList []string `json:"to_wxid_list"`
	Value      int      `json:"value"`
}

type ChatroomAdminOperateResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
	} `json:"data"`
}

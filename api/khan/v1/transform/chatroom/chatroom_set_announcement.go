package chatroom

type ChatroomSetAnnouncementRequest struct {
	Appid   string `json:"appid"`
	Content string `json:"content"`
	GroupId string `json:"group_id"`
}

type ChatroomSetAnnouncementResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		NextTime                  int   `json:"NextTime"`
		Selector                  int64 `json:"Selector"`
		BlueToothBroadCastContent struct {
			ILen int `json:"iLen"`
		} `json:"BlueToothBroadCastContent"`
	} `json:"data"`
}

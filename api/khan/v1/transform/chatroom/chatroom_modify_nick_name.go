package chatroom

type ChatroomModifyNickNameRequest struct {
	Appid    string `json:"appid"`
	GroupId  string `json:"group_id"`
	NickName string `json:"nick_name"`
}

type ChatroomModifyNickNameResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		Ret      int `json:"ret"`
		OplogRet struct {
			Count int    `json:"count"`
			Ret   string `json:"ret"`
		} `json:"oplogRet"`
	} `json:"data"`
}

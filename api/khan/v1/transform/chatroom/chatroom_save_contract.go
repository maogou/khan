package chatroom

type ChatroomMoveContractRequest struct {
	Appid   string `json:"appid"`
	GroupId string `json:"group_id"`
	Value   int    `json:"value"`
}

type ChatroomMoveContractResponse struct {
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

package chatroom

/*
*
如果group_id 是以@chatroom结尾则 true=> 0 false=>1
否则 true=>1 false=>0
*/
type ChatroomSetSilenceRequest struct {
	Appid   string `json:"appid"`
	GroupId string `json:"group_id"`
	Value   int    `json:"value"`
}

type ChatroomSetSilenceResponse struct {
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

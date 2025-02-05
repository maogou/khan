package transform

/*
* 如果group_id以@chatroom结尾 则设置置顶为 2050  不置顶为 2
* 如果group_id不是以@chatroom结尾 则设置置顶为 34823    不置顶为 32775
 */
type GroupTopPingRequest struct {
	Appid   string `json:"appid"`
	GroupId string `json:"group_id"`
	Value   int    `json:"value"` //2050
}

type GroupTopPingResponse struct {
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

package chatroom

type ChatroomAddMemberAsFriendRequest struct {
	Appid   string `json:"appid"`
	Content string `json:"content"`
	GroupId string `json:"group_id"`
	ToWxid  string `json:"to_wxid"`
}

type ChatroomAddMemberAsFriendResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		V1 string `json:"v1"`
		Id string `json:"id"`
	} `json:"data"`
}

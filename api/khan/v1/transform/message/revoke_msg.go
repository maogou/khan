package message

type RevokeMsgRequest struct {
	Appid      string `json:"appid"`
	CreateTime int64  `json:"create_time"`
	NewMsgId   int64  `json:"new_msg_id"`
	ToWxid     string `json:"to_wxid"`
}

type RevokeMsgResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		IsysWording string `json:"isysWording"`
	} `json:"data"`
}

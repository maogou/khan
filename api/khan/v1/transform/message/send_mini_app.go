package message

type SendMiniAppRequest struct {
	Appid    string `json:"appid"`
	FileLink string `json:"file_link"`
	ToWxid   string `json:"to_wxid"`
	Xml      string `json:"xml"`
}

type SendMiniAppResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		FromUserName string `json:"fromUserName"`
		ToUserName   string `json:"toUserName"`
		MsgId        int    `json:"msgId"`
		ClientMsgId  string `json:"clientMsgId"`
		CreateTime   int    `json:"createTime"`
		Type         int    `json:"type"`
		NewMsgId     int64  `json:"newMsgId"`
		MsgSource    string `json:"msgSource"`
		ActionFlag   int    `json:"actionFlag"`
	} `json:"data"`
}

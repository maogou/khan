package transform

type PostLinkRequest struct {
	Appid  string `json:"appid"`
	Base64 string `json:"base64"`
	Des    string `json:"des"`
	Title  string `json:"title"`
	ToWxid string `json:"to_wxid"`
	Url    string `json:"url"`
}

type PostLinkResponse struct {
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

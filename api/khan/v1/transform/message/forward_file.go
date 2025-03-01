package message

type ForwardFileRequest struct {
	Aeskey       string `json:"aeskey"`
	Appid        string `json:"appid"`
	Cdnattachurl string `json:"cdnattachurl"`
	Md5          string `json:"md5"`
	Title        string `json:"title"`
	ToWxid       string `json:"to_wxid"`
	Totallen     int    `json:"totallen"`
}

type ForwardFileResponse struct {
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
		Aeskey       string `json:"aeskey"`
		MsgSource    string `json:"msgSource"`
		ActionFlag   int    `json:"actionFlag"`
	} `json:"data"`
}

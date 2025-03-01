package message

type ForwardVideoRequest struct {
	Appid  string `json:"appid"`
	ToWxid string `json:"to_wxid"`
	Xml    string `json:"xml"`
}

type ForwardVideoResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		ClientMsgId   string `json:"clientMsgId"`
		MsgId         int    `json:"msgId"`
		ThumbStartPos int    `json:"thumbStartPos"`
		VideoStartPos int    `json:"videoStartPos"`
		NewMsgId      int64  `json:"newMsgId"`
		Aeskey        string `json:"aeskey"`
		MsgSource     string `json:"msgSource"`
		ActionFlag    int    `json:"actionFlag"`
	} `json:"data"`
}

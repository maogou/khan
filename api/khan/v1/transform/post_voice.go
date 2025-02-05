package transform

type PostVoiceRequest struct {
	Appid    string `json:"appid"`
	FileLink string `json:"file_link"`
	Second   int    `json:"second"`
	ToWxid   string `json:"to_wxid"`
}

type PostVoiceResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		FromUserName string `json:"FromUserName"`
		ToUserName   string `json:"ToUserName"`
		Offset       int    `json:"Offset"`
		Length       int    `json:"Length"`
		CreateTime   int    `json:"CreateTime"`
		ClientMsgId  string `json:"ClientMsgId"`
		MsgId        int    `json:"MsgId"`
		VoiceLength  int    `json:"VoiceLength"`
		EndFlag      int    `json:"EndFlag"`
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		CancelFlag int   `json:"CancelFlag"`
		NewMsgId   int64 `json:"NewMsgId"`
	} `json:"data"`
}

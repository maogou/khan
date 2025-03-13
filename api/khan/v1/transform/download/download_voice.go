package download

type DownloadVoiceRequest struct {
	Appid        string `json:"appid"`
	Bufid        string `json:"bufid"`
	FromUserName string `json:"from_user_name"`
	Length       int    `json:"length"`
	MsgId        string `json:"msg_id"`
}

type DownloadVoiceResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		MsgId       int    `json:"msgId"`
		Offset      int    `json:"offset"`
		Length      int    `json:"length"`
		VoiceLength int    `json:"voiceLength"`
		ClientMsgId string `json:"clientMsgId"`
		Data        struct {
			ILen   int    `json:"iLen"`
			Buffer string `json:"buffer"`
		} `json:"data"`
		EndFlag      int `json:"endFlag"`
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		CancelFlag int   `json:"cancelFlag"`
		NewMsgId   int64 `json:"newMsgId"`
	} `json:"data"`
}

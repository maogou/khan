package message

type ForwardImageRequest struct {
	Appid  string `json:"appid"`
	ToWxid string `json:"to_wxid"`
	Xml    string `json:"xml"`
}

type ForwardImageResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		Msgid       int `json:"Msgid"`
		ClientImgId struct {
			String string `json:"string"`
		} `json:"ClientImgId"`
		FromUserName struct {
			String string `json:"string"`
		} `json:"FromUserName"`
		ToUserName struct {
			String string `json:"string"`
		} `json:"ToUserName"`
		TotalLen   int    `json:"TotalLen"`
		StartPos   int    `json:"StartPos"`
		DataLen    int    `json:"DataLen"`
		CreateTime int    `json:"CreateTime"`
		Newmsgid   int64  `json:"Newmsgid"`
		Aeskey     string `json:"Aeskey"`
		Fileid     string `json:"Fileid"`
		MsgSource  string `json:"MsgSource"`
		ActionFlag int    `json:"actionFlag"`
	} `json:"data"`
}

package v1

type ToWxid struct {
	Content string `json:"content"`
	MsgType int    `json:"msg_type"`
	ToWxid  string `json:"to_wxid"`
}

type PostTextRequest struct {
	Appid      string   `json:"appid"`
	ToWxidList []ToWxid `json:"to_wxid_list"`
}

type PostTextResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		Count int `json:"Count"`
		List  []struct {
			Ret        int `json:"Ret"`
			ToUsetName struct {
				String string `json:"string"`
			} `json:"ToUsetName"`
			MsgId          int   `json:"MsgId"`
			NewClientMsgid int   `json:"NewClientMsgid"`
			Createtime     int   `json:"Createtime"`
			Servertime     int   `json:"servertime"`
			Type           int   `json:"Type"`
			NewMsgId       int64 `json:"NewMsgId"`
		} `json:"List"`
		NoKnow int `json:"NoKnow"`
	} `json:"data"`
}

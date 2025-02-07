package transform

type ChatroomQrcodeRequest struct {
	AppId   string `json:"appid"`
	GroupId string `json:"group_id"`
}

type ChatroomQrcodeResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		Qrcode struct {
			ILen   int    `json:"iLen"`
			Buffer string `json:"buffer"`
		} `json:"qrcode"`
		RevokeQrcodeWording string `json:"revokeQrcodeWording"`
	} `json:"data"`
}

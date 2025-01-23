package transform

type PersonalQrcodeRequest struct {
	AppId   string `json:"appId" binding:"required"`
	ProxyIp string `json:"proxyIp"`
}

type PersonalQrcodResponse struct {
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
	} `json:"data"`
}

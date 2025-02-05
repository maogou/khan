package transform

type GetLoginQrCodeRequest struct {
	AppId string `json:"appid"`
}

type GetLoginQrCodeResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		Base64 string `json:"base64"`
		Url    string `json:"url"`
		Uuid   string `json:"uuid"`
		Nkey   string `json:"nkey"`
		Id     string `json:"id"`
	} `json:"data"`
}

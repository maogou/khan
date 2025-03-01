package login

type LogoutRequest struct {
	Appid string `json:"appid"`
}

type LogoutResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
	} `json:"data"`
}

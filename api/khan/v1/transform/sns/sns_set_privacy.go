package sns

type SnsSetPrivacyRequest struct {
	AppId string `json:"appid"`
	Id    string `json:"id"`
	Type  int    `json:"type"`
}

type SnsSetPrivacyReponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		OpCount int `json:"opCount"`
	} `json:"data"`
}

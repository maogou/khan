package transform

type PersonalPrivacySettingRequest struct {
	Appid string `json:"appid"`
	Id    int    `json:"id"`
	Value int    `json:"value"`
}

type PersonalPrivacySettingResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		Ret      int `json:"ret"`
		OplogRet struct {
			Count int    `json:"count"`
			Ret   string `json:"ret"`
		} `json:"oplogRet"`
	} `json:"data"`
}

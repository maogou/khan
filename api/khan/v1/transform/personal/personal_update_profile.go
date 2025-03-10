package personal

type PersonalUpdateProfileRequest struct {
	Appid     string `json:"appid"`
	City      string `json:"city"`
	Country   string `json:"country"`
	NickName  string `json:"nick_name"`
	Province  string `json:"province"`
	Sex       int    `json:"sex"`
	Signature string `json:"signature"`
}

type PersonalUpdateProfileResponse struct {
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

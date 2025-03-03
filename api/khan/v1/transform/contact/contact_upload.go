package contact

type ContactUploadMobileRequest struct {
	Appid   string   `json:"appid"`
	Opcode  int      `json:"opcode"`
	PhoneNo []string `json:"phone_no"`
}

type ContactUploadMobileResponse struct {
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

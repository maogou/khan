package transform

type LabelDeleteRequest struct {
	Appid string `json:"appid"`
	Id    string `json:"id"`
}

type LabelDeleteResponse struct {
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

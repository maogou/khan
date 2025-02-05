package transform

type LabelModifyRequest struct {
	Appid      string   `json:"appid"`
	Id         string   `json:"id"`
	ToWxidList []string `json:"to_wxid_list"`
}

type LabelModifyResponse struct {
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

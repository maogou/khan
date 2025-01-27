package transform

type LabelAddRequest struct {
	Appid string `json:"appid"`
	Name  string `json:"name"`
}

type LabelAddResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		LabelCount    int `json:"LabelCount"`
		LabelPairList struct {
			LabelName string `json:"labelName"`
			LabelID   int    `json:"labelID"`
		} `json:"LabelPairList"`
	} `json:"data"`
}

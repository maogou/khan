package label

type LabelListRequest struct {
	Appid string `json:"appid"`
}

type LabelListResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		LabelCount    int `json:"labelCount"`
		LabelPairList []struct {
			LabelName string `json:"labelName"`
			LabelID   int    `json:"labelID"`
		} `json:"labelPairList"`
	} `json:"data"`
}

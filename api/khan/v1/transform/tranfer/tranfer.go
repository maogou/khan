package tranfer

type TranferBatchHeartBeatItem struct {
	Appid string `json:"appid"`
}
type TranferBatchHeartBeatRequest struct {
	List []TranferBatchHeartBeatItem `json:"list"`
}

type TranferBatchHeartBeatResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		List []struct {
			Appid       string `json:"appid"`
			ServiceInfo struct {
				BaseResponse struct {
					Ret    int `json:"ret"`
					ErrMsg struct {
					} `json:"errMsg"`
				} `json:"BaseResponse"`
				NextTime                  int   `json:"NextTime"`
				Selector                  int64 `json:"Selector"`
				BlueToothBroadCastContent struct {
					ILen int `json:"iLen"`
				} `json:"BlueToothBroadCastContent"`
			} `json:"service_info"`
		} `json:"list"`
	} `json:"data"`
}

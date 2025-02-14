package trade

type CollectMoneyRequest struct {
	AppId string `json:"appid"`
	ToWid string `json:"to_wid"`
	Xml   string `json:"xml"`
}

type CollectMoneyResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
				String string `json:"string"`
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		RetText struct {
			ILen   int    `json:"iLen"`
			Buffer string `json:"buffer"`
		} `json:"retText"`
		PlatRet   int    `json:"platRet"`
		PlatMsg   string `json:"platMsg"`
		CgiCmdid  int    `json:"cgiCmdid"`
		ErrorType int    `json:"errorType"`
		ErrorMsg  string `json:"errorMsg"`
	} `json:"data"`
}

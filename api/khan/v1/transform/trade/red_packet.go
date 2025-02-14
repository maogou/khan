package trade

type RedPacketRequest struct {
	AppId   string `json:"appid"`
	GroupId string `json:"group_id"`
	Xml     string `json:"xml"`
}

type RedPacketResponse struct {
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
		CgiCmdid  int    `json:"cgiCmdid"`
		ErrorType int    `json:"errorType"`
		ErrorMsg  string `json:"errorMsg"`
	} `json:"data"`
}

package sns

type SnsDeleteCommentRequest struct {
	AppId     string `json:"appid"`
	Id        string `json:"id"`
	Type      int    `json:"type"`
	CommnetId int    `json:"commnet_id"`
}

type SnsDeleteCommentResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		OpCount int `json:"opCount"`
	} `json:"data"`
}

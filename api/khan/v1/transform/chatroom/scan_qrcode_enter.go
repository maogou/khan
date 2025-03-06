package chatroom

/*
*
{"appid":"wx_yehMbnKBtN80heOkIFvPb","config":{"type":0,"url":"https://weixin.qq.com/g/AwYAAFaT9qeQLWEs8xq4cjHIvcr3q1ND10VAyMO0rHCWoO4PsgQ5spyABQSFdhpo"},"scene":10001}
*/
type ChatroomScanQrcodeEnterRequest struct {
	Appid  string                        `json:"appid"`
	Config ChatroomScanQrcodeEnterConfig `json:"config"`
	Scene  int                           `json:"scene"`
}

type ChatroomScanQrcodeEnterConfig struct {
	Type int    `json:"type"`
	Url  string `json:"url"`
}

type ChatroomScanQrcodeEnterResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		Scene  int `json:"Scene"`
		Result struct {
			Qid       string `json:"qid"`
			GroupName string `json:"group_name"`
			Msg       string `json:"msg"`
		} `json:"Result"`
	} `json:"data"`
}

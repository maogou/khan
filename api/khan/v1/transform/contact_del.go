package transform

/*
*

{"appid":"wx_WijJeN12eAvakvBXkh7m3","to_wxid":"xingmaogou"}
*/
type ContactDelRequest struct {
	Appid  string `json:"appid"`
	ToWxid string `json:"to_wxid"`
}

/*
*
{"ret":0,"msg":"success","data":{"ret":0,"oplogRet":{"count":1,"ret":"AA=="}}}
*/
type ContactDelResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		Ret      int `json:"ret"`
		OplogRet struct {
			Count int    `json:"count"`
			Ret   string `json:"ret"`
		} `json:"oplogRet"`
	} `json:"data"`
}

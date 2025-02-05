package transform

/**
{"appid":"wx_WijJeN12eAvakvBXkh7m3","config":{"status":1,"to_wxid":"xingmaogou"},"func_id":72}
{"appid":"wx_WijJeN12eAvakvBXkh7m3","config":{"status":2,"to_wxid":"xingmaogou"},"func_id":72}
*/

type ContactSetOnlyChatConfig struct {
	Status int    `json:"status"`
	ToWxid string `json:"to_wxid"`
}
type ContactSetOnlyChatRequest struct {
	Appid  string                   `json:"appid"`
	Config ContactSetOnlyChatConfig `json:"config"`
	FuncId int                      `json:"func_id"`
}

/*
*
{"ret":0,"msg":"success","data":{"ret":0,"oplogRet":{"count":1,"ret":"AA=="}}}
*/
type ContactSetOnlyChatResponse struct {
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

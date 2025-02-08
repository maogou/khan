package contact

/*
*
{"appid":"wx_WijJeN12eAvakvBXkh7m3","content":"hello","opcode":2,"scene":3,"v_3":"v3_020b3826fd03010000000000520c10334fe6b3000000501ea9a3dba12f95f6b60a0536a1adb6a45ca29c79e8efeb399e19ac72c3a7e26f7dd4407b1c3246692f4e5998ca85e69272f1e1a3138e18465b28cc3c@stranger","v_4":"v4_000b708f0b040000010000000000e8a963c43e1b21ec31f429d995671000000050ded0b020927e3c97896a09d47e6e9ecbd68f743b4377543f8a336c4dc06a39d3fddea4e84da8d75311b8eb2aa80e0c9168727afad98857fdda60227655b8bf49c6b5eefc2242fc4120e47299c5a80e1e6f93bbc7c6444d51@stranger"}
*/
type ContactAddRequest struct {
	Appid   string `json:"appid"`
	Content string `json:"content"`
	Opcode  int    `json:"opcode"`
	Scene   int    `json:"scene"`
	V3      string `json:"v_3"`
	V4      string `json:"v_4"`
}

/*
*
{"ret":0,"msg":"success","data":{"BaseResponse":{"ret":0,"errMsg":{}},"v1":"v3_020b3826fd03010000000000520c10334fe6b3000000501ea9a3dba12f95f6b60a0536a1adb6a45ca29c79e8efeb399e19ac72c3a7e26f7dd4407b1c3246692f4e5998ca85e69272f1e1a3138e18465b28cc3c@stranger","id":"f8a0171c-1436-4f04-946d-18778d36b3e4"}}
*/
type ContactAddResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		V1 string `json:"v1"`
		Id string `json:"id"`
	} `json:"data"`
}

package message

/*
*
{"appid":"wx_WijJeN12eAvakvBXkh7m3","md5":"4cc7540a85b5b6cf4ba14e9f4ae08b7c","to_wxid_list":["xingmaogou"],"total_len":102357}
*/
type PostEmojiRequest struct {
	Appid      string   `json:"appid"`
	Md5        string   `json:"md5"`
	ToWxidList []string `json:"to_wxid_list"`
	TotalLen   int      `json:"total_len"`
}

/**
{"ret":0,"msg":"success","data":{"BaseResponse":{"ret":0,"errMsg":{}},"emojiItemCount":1,"emojiItem":[{"ret":0,"startPos":102357,"totalLen":102357,"md5":"4cc7540a85b5b6cf4ba14e9f4ae08b7c","msgId":867477064,"newMsgId":6988814996112051511}],"actionFlag":0}}
*/

type PostEmojiResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		EmojiItemCount int `json:"emojiItemCount"`
		EmojiItem      []struct {
			Ret      int    `json:"ret"`
			StartPos int    `json:"startPos"`
			TotalLen int    `json:"totalLen"`
			Md5      string `json:"md5"`
			MsgId    int    `json:"msgId"`
			NewMsgId int64  `json:"newMsgId"`
		} `json:"emojiItem"`
		ActionFlag int `json:"actionFlag"`
	} `json:"data"`
}

package contact

/*
*
{"appid":"wx_WijJeN12eAvakvBXkh7m3","current_chat_room_contact_seq":0,"current_wxcontact_seq":0}
*/
type ContactListRequest struct {
	Appid                     string `json:"appid"`
	CurrentChatRoomContactSeq int    `json:"current_chat_room_contact_seq"`
	CurrentWxcontactSeq       int    `json:"current_wxcontact_seq"`
}

/*
*
{"ret":0,"msg":"success","data":{"BaseResponse":{"ret":0,"errMsg":{"string":""}},"CurrentWxcontactSeq":867477143,"CurrentChatRoomContactSeq":0,"CountinueFlag":0,"ContactUsernameList":["fmessage","medianote","floatbottle","gh_3dfda90e39d6","weixin","filehelper","gh_440797b77f7c","xingmaogou","gh_2b5b65dcbe47","gh_b263da0eb84a","gh_62b365c2817e","gh_44c3853758a7","gh_2cbe88b0b858","gh_594faf5242b6","gh_4b0824b467c6","gh_e6abc2a1ac85","gh_b6614cb49a77","gh_f4580c01069c","gh_4985c77f7fde","wxid_4pusa1tdrs8r12","wxid_500zr8lgqpmv22","wxid_lq2fbfji41lm22","wxid_mfq8hjgz8rd822","gh_219825dd9d7c","qq609654215","gh_187807238e2c","wxid_i9lwn4g45rrb12","gh_e5e0d678bd65","gh_a4c2093e1a9c","wxid_pzd260t9g79h22","gh_75dea2d6c71f","wxid_2sgby07xtcun22","wxid_qp688nulis2l21","gh_e10fff543289","wxid_0z99psrzusvf22","wxid_zngcn5z2gw2i12","wxid_wmvbjpdigha322","forgather"]}}
*/
type ContactListResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
				String string `json:"string"`
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		CurrentWxcontactSeq       int      `json:"CurrentWxcontactSeq"`
		CurrentChatRoomContactSeq int      `json:"CurrentChatRoomContactSeq"`
		CountinueFlag             int      `json:"CountinueFlag"`
		ContactUsernameList       []string `json:"ContactUsernameList"`
	} `json:"data"`
}

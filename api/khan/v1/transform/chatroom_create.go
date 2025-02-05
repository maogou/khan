package transform

type ChatroomCreateRequest struct {
	Appid      string   `json:"appid"`
	ToWxidList []string `json:"to_wxid_list"`
}

type ChatroomCreateResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
				String string `json:"string"`
			} `json:"errMsg"`
		} `json:"baseResponse"`
		Topic struct {
		} `json:"Topic"`
		Pyinitial struct {
		} `json:"Pyinitial"`
		QuanPin struct {
		} `json:"QuanPin"`
		MemberCount int `json:"MemberCount"`
		MemberLis   []struct {
			MemberName struct {
				String string `json:"string"`
			} `json:"MemberName"`
			MemberStatus int `json:"MemberStatus"`
			NickName     struct {
				String string `json:"string"`
			} `json:"NickName"`
			PYInitial struct {
				String string `json:"string"`
			} `json:"PYInitial"`
			QuanPin struct {
				String string `json:"string"`
			} `json:"QuanPin"`
			Sex    int `json:"Sex"`
			Remark struct {
				String string `json:"string,omitempty"`
			} `json:"Remark"`
			RemarkPyinitial struct {
				String string `json:"string,omitempty"`
			} `json:"RemarkPyinitial"`
			RemarkQuanPin struct {
				String string `json:"string,omitempty"`
			} `json:"RemarkQuanPin"`
			ContactType  int    `json:"ContactType"`
			Province     string `json:"Province"`
			City         string `json:"City"`
			Signature    string `json:"Signature"`
			PersonalCard int    `json:"PersonalCard"`
			VerifyFlag   int    `json:"VerifyFlag"`
			Country      string `json:"Country"`
		} `json:"MemberLis"`
		ChatRoomName struct {
			String string `json:"string"`
		} `json:"ChatRoomName"`
		ImgBuf struct {
			ILen   int    `json:"iLen"`
			Buffer string `json:"buffer"`
		} `json:"ImgBuf"`
	} `json:"data"`
}

package v1

type PostTextRequest struct {
	AppId   string `json:"appId" validate:"required"`
	ToWxid  string `json:"toWxid" validate:"required"`  //好友/群的ID
	Ats     string `json:"ats" validate:"omitempty"`    //@的好友，多个英文逗号分隔。群主或管理员@全部的人，则填写'notify@all' 在群内发送消息@某人时，content中需包含@xxx
	Content string `json:"content" validate:"required"` //文本消息
}

type PostTextResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		Count int `json:"Count"`
		List  []struct {
			Ret        int `json:"Ret"`
			ToUsetName struct {
				String string `json:"string"`
			} `json:"ToUsetName"`
			MsgId          int   `json:"MsgId"`
			NewClientMsgid int64 `json:"NewClientMsgid"`
			Createtime     int   `json:"Createtime"`
			Servertime     int   `json:"servertime"`
			Type           int   `json:"Type"`
			NewMsgId       int64 `json:"NewMsgId"`
		} `json:"List"`
		NoKnow int `json:"NoKnow"`
	} `json:"data"`
}

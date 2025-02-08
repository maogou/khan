package message

type PostFileConfig struct {
	FileName string `json:"file_name"`
}

type PostFileRequest struct {
	Appid    string         `json:"appid"`
	Config   PostFileConfig `json:"config"`
	FileLink string         `json:"file_link"`
	ToWxid   string         `json:"to_wxid"`
}

type PostFileResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		UploadStatus struct {
			BaseResponse struct {
				Ret    int `json:"ret"`
				ErrMsg struct {
					String string `json:"string"`
				} `json:"errMsg"`
			} `json:"baseResponse"`
			UploadId           string `json:"uploadId"`
			FileId             string `json:"fileId"`
			AesKey             string `json:"aesKey"`
			ImgInfo            string `json:"imgInfo"`
			CurrentPartIdlimit int    `json:"currentPartIdlimit"`
		} `json:"upload_status"`
		SendMsgStatus struct {
			BaseResponse struct {
				Ret    int `json:"ret"`
				ErrMsg struct {
				} `json:"errMsg"`
			} `json:"BaseResponse"`
			AppId        string `json:"appId"`
			FromUserName string `json:"fromUserName"`
			ToUserName   string `json:"toUserName"`
			MsgId        int    `json:"msgId"`
			ClientMsgId  string `json:"clientMsgId"`
			CreateTime   int    `json:"createTime"`
			Type         int    `json:"type"`
			NewMsgId     int64  `json:"newMsgId"`
			Aeskey       string `json:"aeskey"`
			MsgSource    string `json:"msgSource"`
			ActionFlag   int    `json:"actionFlag"`
		} `json:"send_msg_status"`
		OtherStatus struct {
			BaseResponse struct {
				Ret    int `json:"ret"`
				ErrMsg struct {
					String string `json:"string"`
				} `json:"errMsg"`
			} `json:"baseResponse"`
			FileUploadToken string `json:"fileUploadToken"`
			NewMsgId        int64  `json:"newMsgId"`
		} `json:"other_status"`
		BasicData struct {
		} `json:"basic_data"`
	} `json:"data"`
}

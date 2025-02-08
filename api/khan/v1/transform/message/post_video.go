package message

type PostVideoConfig struct {
	VideoPlayLength int    `json:"video_play_length"`
	VideoThumbnail  string `json:"video_thumbnail"`
}
type PostVideoRequest struct {
	Appid    string          `json:"appid"`
	Config   PostVideoConfig `json:"config"`
	FileLink string          `json:"file_link"`
	ToWxid   string          `json:"to_wxid"`
}

type PostVideoResponse struct {
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
			ClientMsgId   string `json:"clientMsgId"`
			MsgId         int    `json:"msgId"`
			ThumbStartPos int    `json:"thumbStartPos"`
			VideoStartPos int    `json:"videoStartPos"`
			NewMsgId      int64  `json:"newMsgId"`
			Aeskey        string `json:"aeskey"`
			MsgSource     string `json:"msgSource"`
			ActionFlag    int    `json:"actionFlag"`
		} `json:"send_msg_status"`
		OtherStatus struct {
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
		} `json:"other_status"`
		BasicData struct {
			Filekey string `json:"filekey"`
		} `json:"basic_data"`
	} `json:"data"`
}

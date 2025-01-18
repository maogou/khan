package transform

type PostImageConfig struct {
	Thumbnail bool `json:"thumbnail" validate:"required"`
}

type PostImage struct {
	AppId    string          `json:"appid" validate:"required"`
	Config   PostImageConfig `json:"config" validate:"required"`
	FileLink string          `json:"file_link" validate:"required"`
	ToWxid   string          `json:"to_wxid" validate:"required"`
}

type PostImageResponse struct {
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
			Msgid       int `json:"Msgid"`
			ClientImgId struct {
				String string `json:"string"`
			} `json:"ClientImgId"`
			FromUserName struct {
				String string `json:"string"`
			} `json:"FromUserName"`
			ToUserName struct {
				String string `json:"string"`
			} `json:"ToUserName"`
			TotalLen   int    `json:"TotalLen"`
			StartPos   int    `json:"StartPos"`
			DataLen    int    `json:"DataLen"`
			CreateTime int    `json:"CreateTime"`
			Newmsgid   int64  `json:"Newmsgid"`
			Aeskey     string `json:"Aeskey"`
			Fileid     string `json:"Fileid"`
			MsgSource  string `json:"MsgSource"`
			ActionFlag int    `json:"actionFlag"`
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
			Length  int    `json:"length"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
			Filekey string `json:"filekey"`
			FileMd5 string `json:"fileMd5"`
		} `json:"basic_data"`
	} `json:"data"`
}

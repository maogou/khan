package sns

type SnsUploadImgConfig struct {
	Thumbnail bool `json:"thumbnail"`
}
type SnsUploadImgRequest struct {
	Appid    string             `json:"appid"`
	Config   SnsUploadImgConfig `json:"config"`
	FileLink string             `json:"file_link"`
}

type SnsUploadImgResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		Ver         int         `json:"Ver"`
		Seq         int         `json:"Seq"`
		RetCode     int         `json:"RetCode"`
		FileKey     string      `json:"FileKey"`
		FileKey2    string      `json:"FileKey2"`
		FileUrl     string      `json:"FileUrl"`
		Thumburl    string      `json:"Thumburl"`
		RecvLen     int         `json:"RecvLen"`
		SKeyResp    int         `json:"SKeyResp"`
		SKeyBuf     interface{} `json:"SKeyBuf"`
		FileID      string      `json:"FileID"`
		FileIDStr   string      `json:"FileIDStr"`
		ExistFlag   int         `json:"ExistFlag"`
		HitType     int         `json:"HitType"`
		RetrySec    int         `json:"RetrySec"`
		IsRetry     int         `json:"IsRetry"`
		IsOverLoad  int         `json:"IsOverLoad"`
		IsGetCDN    int         `json:"IsGetCDN"`
		XClientIP   string      `json:"XClientIP"`
		AesKey      string      `json:"AesKey"`
		FileMd5     string      `json:"FileMd5"`
		FileMd52    string      `json:"FileMd52"`
		Mp4Identify string      `json:"Mp4Identify"`
		ImgInfo     struct {
			Length int    `json:"Length"`
			Width  int    `json:"Width"`
			Height int    `json:"Height"`
			MD5    string `json:"MD5"`
		} `json:"ImgInfo"`
		Thumb struct {
			Length int    `json:"Length"`
			Width  int    `json:"Width"`
			Height int    `json:"Height"`
			MD5    string `json:"MD5"`
		} `json:"Thumb"`
		Midimg struct {
			Length int    `json:"Length"`
			Width  int    `json:"Width"`
			Height int    `json:"Height"`
			MD5    string `json:"MD5"`
		} `json:"Midimg"`
		Rsw struct {
			Length int    `json:"Length"`
			Width  int    `json:"Width"`
			Height int    `json:"Height"`
			MD5    string `json:"MD5"`
		} `json:"Rsw"`
	} `json:"data"`
}

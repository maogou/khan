package download

type DownloadImgRequest struct {
	AesKey    string `json:"aes_key"`
	Appid     string `json:"appid"`
	FileType  int    `json:"file_type"`
	FileUrl   string `json:"file_url"`
	Totalsize int    `json:"totalsize"`
}

type DownloadImgResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		Ver             int    `json:"Ver"`
		Seq             int    `json:"Seq"`
		VideoFormat     int    `json:"VideoFormat"`
		RspPicFormat    int    `json:"RspPicFormat"`
		RangeStart      int    `json:"RangeStart"`
		RangeEnd        int    `json:"RangeEnd"`
		TotalSize       int    `json:"TotalSize"`
		SrcSize         int    `json:"SrcSize"`
		RetCode         int    `json:"RetCode"`
		SubStituteFType int    `json:"SubStituteFType"`
		RetrySec        int    `json:"RetrySec"`
		IsRetry         int    `json:"IsRetry"`
		IsOverLoad      int    `json:"IsOverLoad"`
		IsGetCdn        int    `json:"IsGetCdn"`
		XClientIP       string `json:"XClientIP"`
		FileData        string `json:"FileData"`
	} `json:"data"`
}

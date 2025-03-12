package download

type DownloadCdnRequest struct {
	Appid     string `json:"appid"`
	AesKey    string `json:"aes_key"`
	FileUrl   string `json:"file_url"`
	Totalsize int    `json:"totalsize"`
	FileType  int    `json:"file_type"`
}

type DownloadCdnResponse struct {
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

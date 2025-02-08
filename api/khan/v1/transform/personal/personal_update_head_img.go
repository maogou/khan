package personal

type PersonalUploadHdHeadImgRequest struct {
	Appid    string `json:"appid"`
	FileLink string `json:"file_link"`
}

type PersonalUploadHdHeadImgResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"baseResponse"`
		TotalLen        int    `json:"TotalLen"`
		StartPos        int    `json:"StartPos"`
		BigHeadImgUrl   string `json:"BigHeadImgUrl"`
		SmallHeadImgUrl string `json:"SmallHeadImgUrl"`
	} `json:"data"`
}

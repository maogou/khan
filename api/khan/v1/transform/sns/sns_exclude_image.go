package sns

type SnsSendExcludeImageRequest struct {
	AppId        string   `json:"appId"`
	AllowUser    []string `json:"allowWxIds"`
	AllowTagId   []string `json:"allowTagIds"`
	AtUser       []string `json:"atWxIds"`
	DisableUser  []string `json:"disableWxIds"`
	DisableTagId []string `json:"disableTagIds"`
	Private      bool     `json:"private"`
	XmlTxt       string   `json:"snsXml"`
}

type SnsSendExcludeImageResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		SnsObject struct {
			Id         float64 `json:"Id"`
			Username   string  `json:"Username"`
			Nickname   string  `json:"Nickname"`
			CreateTime int64   `json:"CreateTime"`
			ObjectDesc struct {
				ILen   int    `json:"iLen"`
				Buffer string `json:"buffer"`
			} `json:"ObjectDesc"`
			LikeFlag             int `json:"LikeFlag"`
			LikeCount            int `json:"LikeCount"`
			LikeUserListCount    int `json:"LikeUserListCount"`
			CommentCount         int `json:"CommentCount"`
			CommentUserListCount int `json:"CommentUserListCount"`
			WithUserCount        int `json:"WithUserCount"`
			WithUserListCount    int `json:"WithUserListCount"`
			ExtFlag              int `json:"ExtFlag"`
			NoChange             int `json:"NoChange"`
			GroupCount           int `json:"GroupCount"`
			IsNotRichText        int `json:"IsNotRichText"`
			ReferId              int `json:"ReferId"`
			BlackListCount       int `json:"BlackListCount"`
			DeleteFlag           int `json:"DeleteFlag"`
			GroupUserCount       int `json:"GroupUserCount"`
			ObjectOperations     []struct {
				ILen int `json:"iLen"`
			} `json:"ObjectOperations"`
			SnsRedEnvelops struct {
				RewardCount int `json:"RewardCount"`
				ReportId    int `json:"ReportId"`
				ReportKey   int `json:"ReportKey"`
				ResourceId  int `json:"ResourceId"`
			} `json:"SnsRedEnvelops"`
			PreDownloadInfo struct {
				PreDownloadPercent int `json:"PreDownloadPercent"`
				PreDownloadNetType int `json:"PreDownloadNetType"`
			} `json:"PreDownloadInfo"`
			WeAppInfo struct {
				AppId    int `json:"AppId"`
				ShowType int `json:"ShowType"`
				RScore   int `json:"RScore"`
			} `json:"WeAppInfo"`
		} `json:"SnsObject"`
		Id string `json:"id"`
	} `json:"data"`
}

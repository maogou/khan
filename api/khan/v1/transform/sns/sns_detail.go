package sns

type SnsDetailRequest struct {
	AppId   string `json:"appid"`
	Decrypt bool   `json:"decrypt"`
	Id      string `json:"id"`
}

type SnsDetailResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"baseResponse"`
		Object struct {
			Id         int64  `json:"Id"`
			Username   string `json:"Username"`
			Nickname   string `json:"Nickname"`
			CreateTime int64  `json:"CreateTime"`
			ObjectDesc struct {
				ILen   int    `json:"iLen"`
				Buffer string `json:"buffer"`
			} `json:"ObjectDesc"`
			LikeFlag          int `json:"LikeFlag"`
			LikeCount         int `json:"LikeCount"`
			LikeUserListCount int `json:"LikeUserListCount"`
			LikeUserList      []struct {
				Username        string `json:"Username"`
				Nickname        string `json:"Nickname"`
				Source          int    `json:"Source"`
				Type            int    `json:"Type"`
				CreateTime      int64  `json:"CreateTime"`
				CommentId       int    `json:"CommentId"`
				ReplyCommentId  int    `json:"ReplyCommentId"`
				IsNotRichText   int    `json:"IsNotRichText"`
				ReplyCommentId2 int    `json:"ReplyCommentId2"`
				CommentId2      int    `json:"CommentId2"`
				DeleteFlag      int    `json:"DeleteFlag"`
				CommentFlag     int    `json:"CommentFlag"`
			} `json:"LikeUserList"`
			CommentCount         int `json:"CommentCount"`
			CommentUserListCount int `json:"CommentUserListCount"`
			CommentUserList      []struct {
				Username        string `json:"Username"`
				Nickname        string `json:"Nickname"`
				Source          int    `json:"Source"`
				Type            int    `json:"Type"`
				Content         string `json:"Content"`
				CreateTime      int64  `json:"CreateTime"`
				CommentId       int    `json:"CommentId"`
				ReplyCommentId  int    `json:"ReplyCommentId"`
				IsNotRichText   int    `json:"IsNotRichText"`
				ReplyCommentId2 int    `json:"ReplyCommentId2"`
				CommentId2      int    `json:"CommentId2"`
				DeleteFlag      int    `json:"DeleteFlag"`
				CommentFlag     int    `json:"CommentFlag"`
			} `json:"CommentUserList"`
			WithUserCount     int `json:"WithUserCount"`
			WithUserListCount int `json:"WithUserListCount"`
			WithUserList      []struct {
				Username        string `json:"Username"`
				Source          int    `json:"Source"`
				Type            int    `json:"Type"`
				CreateTime      int64  `json:"CreateTime"`
				CommentId       int    `json:"CommentId"`
				ReplyCommentId  int    `json:"ReplyCommentId"`
				IsNotRichText   int    `json:"IsNotRichText"`
				ReplyCommentId2 int    `json:"ReplyCommentId2"`
				CommentId2      int    `json:"CommentId2"`
				DeleteFlag      int    `json:"DeleteFlag"`
				CommentFlag     int    `json:"CommentFlag"`
			} `json:"WithUserList"`
			ExtFlag          int `json:"ExtFlag"`
			NoChange         int `json:"NoChange"`
			GroupCount       int `json:"GroupCount"`
			IsNotRichText    int `json:"IsNotRichText"`
			ReferId          int `json:"ReferId"`
			BlackListCount   int `json:"BlackListCount"`
			DeleteFlag       int `json:"DeleteFlag"`
			GroupUserCount   int `json:"GroupUserCount"`
			ObjectOperations []struct {
				ILen   int    `json:"iLen"`
				Buffer string `json:"buffer"`
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
		} `json:"object"`
	} `json:"data"`
}

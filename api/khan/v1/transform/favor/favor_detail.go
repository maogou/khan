package favor

type FavorDetailRequest struct {
	AppId string `json:"appid"`
	FavId int    `json:"fav_id"`
}

type FavorDetailResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		Count      int `json:"Count"`
		ObjectList []struct {
			FavId      int    `json:"FavId"`
			Status     int    `json:"Status"`
			Object     string `json:"Object"`
			Flag       int    `json:"Flag"`
			UpdateTime int    `json:"UpdateTime"`
			UpdateSeq  int    `json:"UpdateSeq"`
		} `json:"ObjectList"`
	} `json:"data"`
}

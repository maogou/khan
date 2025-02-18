package favor

type FavorDeleteRequest struct {
	AppId string `json:"appid"`
	FavId int    `json:"fav_id"`
}

type FavorDeleteResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"BaseResponse"`
		Count int `json:"Count"`
		List  []struct {
			Ret   int `json:"Ret"`
			FavId int `json:"FavId"`
		} `json:"List"`
	} `json:"data"`
}

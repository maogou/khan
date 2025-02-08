package label

type FavorSyncRequest struct {
	AppId   string `json:"appid"`
	SyncKey string `json:"sync_key"`
}

type FavorSyncResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		Ret  int `json:"Ret"`
		List []struct {
			FavId      int `json:"FavId"`
			Type       int `json:"Type"`
			Flag       int `json:"Flag"`
			UpdateTime int `json:"UpdateTime"`
			UpdateSeq  int `json:"UpdateSeq"`
		} `json:"List"`
		KeyBuf struct {
			ILen   int    `json:"iLen"`
			Buffer string `json:"buffer"`
		} `json:"KeyBuf"`
	} `json:"data"`
}

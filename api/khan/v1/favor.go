package v1

type FavorSyncRequest struct {
	AppId   string `json:"appId" binding:"required"`
	SyncKey string `json:"syncKey"`
}

type FavorSyncData struct {
	Data struct {
		SyncKey string          `json:"syncKey"`
		List    []FavorSyncItem `json:"list"`
	} `json:"data"`
}

type FavorSyncItem struct {
	FavId      int `json:"favId"`
	Type       int `json:"type"`
	Flag       int `json:"flag"`
	UpdateTime int `json:"updateTime"`
}

type FavorSyncResponse struct {
	ResponseCommon
	FavorSyncData
}

type FavorDetailRequest struct {
	AppId string `json:"appId" binding:"required"`
	FavId int    `json:"favId" binding:"required"`
}

type FavorDetailData struct {
	FavId      int    `json:"favId"`
	Status     int    `json:"status"`
	Flag       int    `json:"flag"`
	UpdateTime int    `json:"updateTime"`
	Content    string `json:"content"`
}

type FavorDetailResponse struct {
	ResponseCommon
	Data FavorDetailData `json:"data"`
}

type FavorDeleteRequest struct {
	AppId string `json:"appId" binding:"required"`
	FavId int    `json:"favId" binding:"required"`
}

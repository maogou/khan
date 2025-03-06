package contact

type ContactGetMobileRequest struct {
	Appid   string   `json:"appid"`
	PhoneNo []string `json:"phone_no"`
}

type ContactGetMobileResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
				String string `json:"string"`
			} `json:"errMsg"`
		} `json:"baseResponse"`
		Count      int `json:"Count"`
		FriendList []struct {
			UserName     string `json:"UserName"`
			Nickname     string `json:"Nickname"`
			MobileMD5    string `json:"MobileMD5"`
			Sex          int    `json:"Sex"`
			Province     string `json:"Province"`
			City         string `json:"City"`
			Signature    string `json:"Signature"`
			PersonalCard int    `json:"PersonalCard"`
			FBInfo       struct {
				Id     int `json:"id"`
				ImgKey int `json:"imgKey"`
			} `json:"FBInfo"`
			AlbumFlag   int `json:"AlbumFlag"`
			AlbumStyle  int `json:"AlbumStyle"`
			SnsUserInfo struct {
				SnsFlag       int `json:"SnsFlag"`
				SnsBgobjectId int `json:"SnsBgobjectId"`
				SnsFlagEx     int `json:"SnsFlagEx"`
			} `json:"SnsUserInfo"`
			Country        string `json:"Country"`
			CustomizedInfo struct {
				BrandFlag int `json:"BrandFlag"`
			} `json:"CustomizedInfo"`
			BigHeadImgUrl   string `json:"BigHeadImgUrl"`
			SmallHeadImgUrl string `json:"SmallHeadImgUrl"`
		} `json:"FriendList"`
	} `json:"data"`
}

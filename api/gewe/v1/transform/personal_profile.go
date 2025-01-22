package transform

type PersonalProfileRequest struct {
	AppId   string `json:"appId" binding:"required"`
	ProxyIp string `json:"proxyIp"`
}

type PersonalProfileResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		BaseResponse struct {
			Ret    int `json:"ret"`
			ErrMsg struct {
			} `json:"errMsg"`
		} `json:"baseResponse"`
		UserInfo struct {
			BitFlag  int `json:"BitFlag"`
			UserName struct {
				String string `json:"string"`
			} `json:"UserName"`
			NickName struct {
				String string `json:"string"`
			} `json:"NickName"`
			BindUin   int `json:"BindUin"`
			BindEmail struct {
			} `json:"BindEmail"`
			BindMobile struct {
			} `json:"BindMobile"`
			Status         int    `json:"Status"`
			ImgLen         int    `json:"ImgLen"`
			Sex            int    `json:"Sex"`
			Province       string `json:"Province"`
			Signature      string `json:"Signature"`
			PersonalCard   int    `json:"PersonalCard"`
			DisturbSetting struct {
				NightSetting int `json:"NightSetting"`
				NightTime    struct {
					BeginTime int `json:"BeginTime"`
					EndTime   int `json:"EndTime"`
				} `json:"NightTime"`
				AllDaySetting int `json:"AllDaySetting"`
				AllDayTim     struct {
					BeginTime int `json:"BeginTime"`
					EndTime   int `json:"EndTime"`
				} `json:"AllDayTim"`
			} `json:"DisturbSetting"`
			PluginFlag   int `json:"PluginFlag"`
			VerifyFlag   int `json:"VerifyFlag"`
			Point        int `json:"Point"`
			Experience   int `json:"Experience"`
			Level        int `json:"Level"`
			LevelLowExp  int `json:"LevelLowExp"`
			LevelHighExp int `json:"LevelHighExp"`
			PluginSwitch int `json:"PluginSwitch"`
			GmailList    struct {
				Count int `json:"Count"`
			} `json:"GmailList"`
			Alias          string `json:"Alias"`
			WeiboFlag      int    `json:"WeiboFlag"`
			FaceBookFlag   int    `json:"FaceBookFlag"`
			FbuserId       int    `json:"FbuserId"`
			AlbumStyle     int    `json:"AlbumStyle"`
			AlbumFlag      int    `json:"AlbumFlag"`
			TxnewsCategory int    `json:"TxnewsCategory"`
			Country        string `json:"Country"`
		} `json:"userInfo"`
		UserInfoExt struct {
			SnsUserInfo struct {
				SnsFlag       int     `json:"SnsFlag"`
				SnsBgimgId    string  `json:"SnsBgimgId"`
				SnsBgobjectId float64 `json:"SnsBgobjectId"`
				SnsFlagEx     int     `json:"SnsFlagEx"`
			} `json:"SnsUserInfo"`
			MyBrandList       string `json:"MyBrandList"`
			BigChatRoomSize   int    `json:"BigChatRoomSize"`
			BigChatRoomQuota  int    `json:"BigChatRoomQuota"`
			BigChatRoomInvite int    `json:"BigChatRoomInvite"`
			BigHeadImgUrl     string `json:"BigHeadImgUrl"`
			SmallHeadImgUrl   string `json:"SmallHeadImgUrl"`
			MainAcctType      int    `json:"MainAcctType"`
			ExtXml            struct {
			} `json:"ExtXml"`
			SafeDeviceList struct {
				Count int `json:"Count"`
				List  []struct {
					Name       string `json:"Name"`
					Uuid       string `json:"Uuid"`
					DeviceType string `json:"DeviceType"`
					CreateTime int    `json:"CreateTime"`
				} `json:"List"`
			} `json:"SafeDeviceList"`
			SafeDevice          int    `json:"SafeDevice"`
			GrayscaleFlag       int    `json:"GrayscaleFlag"`
			RegCountry          string `json:"RegCountry"`
			LinkedinContactItem struct {
			} `json:"LinkedinContactItem"`
			PatternLockInfo struct {
				PatternVersion int `json:"PatternVersion"`
				Sign           struct {
					ILen   int    `json:"iLen"`
					Buffer string `json:"buffer"`
				} `json:"Sign"`
				LockStatus int `json:"LockStatus"`
			} `json:"PatternLockInfo"`
			PayWalletType int     `json:"PayWalletType"`
			WalletRegion  int     `json:"WalletRegion"`
			ExtStatus     float64 `json:"ExtStatus"`
			UserStatus    int     `json:"UserStatus"`
		} `json:"userInfoExt"`
	} `json:"data"`
}
